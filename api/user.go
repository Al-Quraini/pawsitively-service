package api

import (
	"database/sql"
	"net/http"
	"time"

	db "github.com/alquraini/pawsitively/db/sqlc"
	"github.com/alquraini/pawsitively/token"
	"github.com/alquraini/pawsitively/util"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type userRespnse struct {
	ID        int64      `json:"id"`
	Email     string     `json:"email"`
	FullName  *string    `json:"full_name"`
	City      *string    `json:"city"`
	State     *string    `json:"state"`
	Country   *string    `json:"country"`
	ImageID   *int64     `json:"image_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func newUserReponse(user db.User) userRespnse {

	return userRespnse{
		ID:        user.ID,
		Email:     user.Email,
		FullName:  util.GetNullableString(user.FullName),
		City:      util.GetNullableString(user.City),
		State:     util.GetNullableString(user.State),
		Country:   util.GetNullableString(user.Country),
		ImageID:   util.GetNullableInt64(user.ImageID),
		CreatedAt: user.CreatedAt,
		UpdatedAt: util.GetNullableTime(user.UpdatedAt),
	}
}

// ****************************** //
// login user   //
// ****************************** //
type loginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string      `json:"access_token"`
	User        userRespnse `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.action.GetUser(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(
		user.ID,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginUserResponse{
		AccessToken: accessToken,
		User:        newUserReponse(user),
	}

	ctx.JSON(http.StatusOK, rsp)
}

// ****************************** //
// register user   //
// ****************************** //
type registerUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (server *Server) registerUser(ctx *gin.Context) {
	var req registerUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Email:          req.Email,
		HashedPassword: hashedPassword,
	}

	user, err := server.action.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(
		user.ID,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginUserResponse{
		AccessToken: accessToken,
		User:        newUserReponse(user),
	}

	ctx.JSON(http.StatusOK, rsp)
}

// ****************************** //
// update user   //
// ****************************** //
type updateUserRequest struct {
	FullName string `json:"full_name"`
	City     string `json:"city"`
	State    string `json:"state"`
	Country  string `json:"country"`
	ImageID  int64  `json:"image_id"`
}

func (server *Server) updateUser(ctx *gin.Context) {
	var req updateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.UpdateUserParams{
		ID:       authPayload.UserID,
		FullName: util.StringToNullString(req.FullName),
		City:     util.StringToNullString(req.City),
		State:    util.StringToNullString(req.State),
		Country:  util.StringToNullString(req.Country),
		ImageID:  util.Int64ToNullInt64(req.ImageID),
	}

	user, err := server.action.UpdateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newUserReponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

// ****************************** //
// delete user   //
// ****************************** //

func (server *Server) deleteUser(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	err := server.action.DeleteUser(ctx, authPayload.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, successResponse("User has been deleted successfully"))
}
