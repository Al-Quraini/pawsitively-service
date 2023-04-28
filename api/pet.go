package api

import (
	"errors"
	"net/http"
	"time"

	db "github.com/alquraini/pawsitively/db/sqlc"
	"github.com/alquraini/pawsitively/token"
	"github.com/alquraini/pawsitively/util"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type petResponse struct {
	ID               int64      `json:"id"`
	Name             string     `json:"name"`
	About            *string    `json:"about"`
	UserID           int64      `json:"user_id"`
	Age              int32      `json:"age"`
	Gender           string     `json:"gender"`
	PetType          string     `json:"pet_type"`
	Breed            *string    `json:"breed"`
	ImageUrl         *string    `json:"image_url"`
	MedicalCondition *string    `json:"medical_condition"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        *time.Time `json:"updated_at"`
}

func createNewPetResponse(pet db.Pet) petResponse {
	return petResponse{
		ID:               pet.ID,
		Name:             pet.Name,
		About:            util.GetNullableString(pet.About),
		UserID:           pet.UserID,
		Age:              pet.Age,
		Gender:           pet.Gender,
		PetType:          pet.PetType,
		Breed:            util.GetNullableString(pet.Breed),
		ImageUrl:         util.GetNullableString(pet.ImageUrl),
		MedicalCondition: util.GetNullableString(pet.MedicalCondition),
		CreatedAt:        pet.CreatedAt,
		UpdatedAt:        util.GetNullableTime(pet.UpdatedAt),
	}
}

// ****************************** //
// create pet   //
// ****************************** //
type createPetRequest struct {
	Name             string `json:"name"`
	About            string `json:"about"`
	Age              int32  `json:"age"`
	Gender           string `json:"gender"`
	PetType          string `json:"pet_type"`
	Breed            string `json:"breed"`
	ImageUrl         string `json:"image_url"`
	MedicalCondition string `json:"medical_condition"`
}

func (server *Server) createNewPet(ctx *gin.Context) {
	var req createPetRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.CreatePetParams{
		Name:             req.Name,
		About:            util.StringToNullString(req.About),
		UserID:           authPayload.UserID,
		Age:              req.Age,
		Gender:           req.Gender,
		PetType:          req.PetType,
		Breed:            util.StringToNullString(req.Breed),
		ImageUrl:         util.StringToNullString(req.ImageUrl),
		MedicalCondition: util.StringToNullString(req.MedicalCondition),
	}

	pet, err := server.action.CreatePet(ctx, arg)
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

	rsp := createNewPetResponse(pet)
	ctx.JSON(http.StatusOK, rsp)
}

type getPetsRequest struct {
	UserID int64 `form:"user_id" binding:"required,min=0"`
}

func (server *Server) getPets(ctx *gin.Context) {
	var req getPetsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	pets, err := server.action.GetPets(ctx, req.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	petResponses := make([]petResponse, len(pets))
	for i, pet := range pets {
		petResponses[i] = createNewPetResponse(pet)
	}
	ctx.JSON(http.StatusOK, petResponses)
}

type getPetByIdRequest struct {
	ID int64 `uri:"id" binding:"required,min=0"`
}

func (server *Server) getPetById(ctx *gin.Context) {
	var req getPetByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	pet, err := server.action.GetPetById(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := createNewPetResponse(pet)
	ctx.JSON(http.StatusOK, rsp)
}

type updatePetRequest struct {
	ID               int64  `json:"id" binding:"required"`
	Name             string `json:"name"`
	About            string `json:"about"`
	Age              int32  `json:"age"`
	UserID           int64  `json:"user_id" binding:"required"`
	Gender           string `json:"gender"`
	PetType          string `json:"pet_type"`
	Breed            string `json:"breed"`
	ImageUrl         string `json:"image_url"`
	MedicalCondition string `json:"medical_condition"`
}

func (server *Server) updatePet(ctx *gin.Context) {
	var req updatePetRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if req.UserID != authPayload.UserID {
		err := errors.New("account doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	arg := db.UpdatePetParams{
		ID:               req.ID,
		Name:             req.Name,
		About:            util.StringToNullString(req.About),
		Age:              req.Age,
		Gender:           req.Gender,
		PetType:          req.PetType,
		Breed:            util.StringToNullString(req.Breed),
		ImageUrl:         util.StringToNullString(req.ImageUrl),
		MedicalCondition: util.StringToNullString(req.MedicalCondition),
	}

	pet, err := server.action.UpdatePet(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := createNewPetResponse(pet)
	ctx.JSON(http.StatusOK, rsp)
}

type deletePetRequest struct {
	ID     int64 `json:"id" binding:"required"`
	UserID int64 `json:"user_id" binding:"required"`
}

func (server *Server) deletePet(ctx *gin.Context) {
	var req deletePetRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if req.UserID != authPayload.UserID {
		err := errors.New("account doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	err := server.action.DeletePet(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, successResponse("Item has been deleted successfully"))
}
