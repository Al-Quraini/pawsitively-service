package api

import (
	"errors"
	"net/http"
	"time"

	db "github.com/alquraini/pawsitively/db/sqlc"
	"github.com/alquraini/pawsitively/token"
	"github.com/alquraini/pawsitively/util"
	"github.com/gin-gonic/gin"
)

type postResponse struct {
	ID         int64      `json:"id"`
	Title      *string    `json:"title"`
	Body       *string    `json:"body"`
	UserID     int64      `json:"user_id"`
	ImageUrl   *string    `json:"image_url"`
	Status     *string    `json:"status"`
	LikesCount int32      `json:"likes_count"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

func createNewPostResponse(post db.Post) postResponse {
	return postResponse{
		ID:         post.ID,
		Title:      util.GetNullableString(post.Title),
		Body:       util.GetNullableString(post.Body),
		UserID:     post.UserID,
		ImageUrl:   util.GetNullableString(post.ImageUrl),
		Status:     util.GetNullableString(post.Status),
		LikesCount: post.LikesCount,
		CreatedAt:  post.CreatedAt,
		UpdatedAt:  util.GetNullableTime(post.UpdatedAt),
	}
}

type createPostRequest struct {
	Title    string `json:"title" binding:"required"`
	Body     string `json:"body"`
	ImageURL string `json:"image_url"`
	Status   string `json:"status"`
}

func (server *Server) createPost(ctx *gin.Context) {
	var req createPostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.CreatePostParams{
		Title:    util.StringToNullString(req.Title),
		Body:     util.StringToNullString(req.Body),
		UserID:   authPayload.UserID,
		ImageUrl: util.StringToNullString(req.ImageURL),
		Status:   util.StringToNullString(req.Status),
	}

	post, err := server.action.CreatePost(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := createNewPostResponse(post)
	ctx.JSON(http.StatusOK, rsp)
}

type getPostRequest struct {
	ID int64 `uri:"id" binding:"required,min=0"`
}

func (server *Server) getPost(ctx *gin.Context) {
	var req getPostRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	post, err := server.action.GetPost(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := createNewPostResponse(post)
	ctx.JSON(http.StatusOK, rsp)
}

type listPostsRequest struct {
	Offset int32 `form:"offset"`
	Limit  int32 `form:"limit" binding:"required,min=5,max=20"`
}

func (server *Server) listPosts(ctx *gin.Context) {
	var req listPostsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListPostsParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	}

	posts, err := server.action.ListPosts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	postsResponses := make([]postResponse, len(posts))
	for i, post := range posts {
		postsResponses[i] = createNewPostResponse(post)
	}
	ctx.JSON(http.StatusOK, postsResponses)
}

type listPostsRequestByUserID struct {
	Offset int32 `form:"offset"`
	Limit  int32 `form:"limit" binding:"required,min=5,max=20"`
}
type listPostsUserIDRequest struct {
	UserID int64 `uri:"user_id" binding:"required,min=0"`
}

func (server *Server) listPostsByUserID(ctx *gin.Context) {
	var userIdReq listPostsUserIDRequest
	if err := ctx.ShouldBindUri(&userIdReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var req listPostsRequestByUserID
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListPostsByUserIDParams{
		UserID: userIdReq.UserID,
		Limit:  req.Limit,
		Offset: req.Offset,
	}

	posts, err := server.action.ListPostsByUserID(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	postsResponses := make([]postResponse, len(posts))
	for i, post := range posts {
		postsResponses[i] = createNewPostResponse(post)
	}
	ctx.JSON(http.StatusOK, postsResponses)
}

type updatePostRequest struct {
	ID       int64  `json:"id"`
	Title    string `json:"title" binding:"required"`
	Body     string `json:"body"`
	UserID   int64  `json:"user_id"`
	ImageURL string `json:"image_url"`
	Status   string `json:"status"`
}

func (server *Server) updatePost(ctx *gin.Context) {
	var req updatePostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if req.UserID != authPayload.UserID {
		err := errors.New("post doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	arg := db.UpdatePostParams{
		ID:       req.ID,
		Title:    util.StringToNullString(req.Title),
		Body:     util.StringToNullString(req.Body),
		ImageUrl: util.StringToNullString(req.ImageURL),
		Status:   util.StringToNullString(req.Status),
	}

	post, err := server.action.UpdatePost(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := createNewPostResponse(post)
	ctx.JSON(http.StatusOK, rsp)
}

type deletePostRequest struct {
	ID     int64 `json:"id" binding:"required"`
	UserID int64 `json:"user_id" binding:"required"`
}

func (server *Server) deletePost(ctx *gin.Context) {
	var req deletePostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if req.UserID != authPayload.UserID {
		err := errors.New("post doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	err := server.action.DeletePost(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, successResponse("Post has been deleted successfully"))
}
