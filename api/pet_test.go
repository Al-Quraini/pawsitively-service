package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mockdb "github.com/alquraini/pawsitively/db/mock"
	db "github.com/alquraini/pawsitively/db/sqlc"
	"github.com/alquraini/pawsitively/token"
	"github.com/alquraini/pawsitively/util"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreatePet(t *testing.T) {
	user, _ := randomUser(t)
	pet := randomPet(t, user)

	testCases := []struct {
		name          string
		body          gin.H
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(action *mockdb.MockAction)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"name":              pet.Name,
				"about":             pet.About.String,
				"age":               pet.Age,
				"gender":            pet.Gender,
				"pet_type":          pet.PetType,
				"breed":             pet.Breed.String,
				"image_url":         pet.ImageUrl.String,
				"medical_condition": pet.MedicalCondition.String,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			},
			buildStubs: func(action *mockdb.MockAction) {
				arg := db.CreatePetParams{
					Name:             pet.Name,
					About:            pet.About,
					Age:              pet.Age,
					UserID:           user.ID,
					Gender:           pet.Gender,
					PetType:          pet.PetType,
					Breed:            pet.Breed,
					ImageUrl:         pet.ImageUrl,
					MedicalCondition: pet.MedicalCondition,
				}
				action.EXPECT().
					CreatePet(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(pet, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "BadRequest",
			body: gin.H{
				"name":              1,
				"about":             pet.About.String,
				"age":               pet.Age,
				"gender":            pet.Gender,
				"pet_type":          pet.PetType,
				"breed":             pet.Breed.String,
				"image_url":         pet.ImageUrl.String,
				"medical_condition": pet.MedicalCondition.String,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			},
			buildStubs: func(action *mockdb.MockAction) {
				arg := db.CreatePetParams{
					Name:             pet.Name,
					About:            pet.About,
					Age:              pet.Age,
					UserID:           user.ID,
					Gender:           pet.Gender,
					PetType:          pet.PetType,
					Breed:            pet.Breed,
					ImageUrl:         pet.ImageUrl,
					MedicalCondition: pet.MedicalCondition,
				}
				action.EXPECT().
					CreatePet(gomock.Any(), gomock.Eq(arg)).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			action := mockdb.NewMockAction(ctrl)
			tc.buildStubs(action)

			server := newTestServer(t, action)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/users/pets/create"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func TestGetPets(t *testing.T) {
	user, _ := randomUser(t)
	var lastPet db.Pet

	n := 5
	pets := make([]db.Pet, n)
	for i := 0; i < n; i++ {
		pets[i] = randomPet(t, user)
		lastPet = pets[i]
	}

	testCases := []struct {
		name          string
		userID        int64
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(action *mockdb.MockAction)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			userID: lastPet.UserID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			},
			buildStubs: func(action *mockdb.MockAction) {
				action.EXPECT().
					GetPets(gomock.Any(), gomock.Eq(lastPet.UserID)).
					Times(1).
					Return(pets, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:   "BadRequest",
			userID: -1,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			},
			buildStubs: func(action *mockdb.MockAction) {
				action.EXPECT().
					GetPets(gomock.Any(), gomock.Eq(-1)).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockAction(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := "/users/pets"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			q := request.URL.Query()
			q.Add("user_id", fmt.Sprintf("%d", tc.userID))
			request.URL.RawQuery = q.Encode()

			// url := fmt.Sprintf("/pets/%d", tc.userID)
			// request, err := http.NewRequest(http.MethodGet, url, nil)
			// require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func TestGetPetById(t *testing.T) {
	user, _ := randomUser(t)
	pet := randomPet(t, user)

	testCases := []struct {
		name          string
		ID            int64
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(action *mockdb.MockAction)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			ID:   pet.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			},
			buildStubs: func(action *mockdb.MockAction) {
				action.EXPECT().
					GetPetById(gomock.Any(), gomock.Eq(pet.ID)).
					Times(1).
					Return(pet, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "BadRequest",
			ID:   -1,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			},
			buildStubs: func(action *mockdb.MockAction) {
				action.EXPECT().
					GetPetById(gomock.Any(), gomock.Eq(-1)).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockAction(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/users/pets/%d", tc.ID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func TestUpdatePet(t *testing.T) {
	user, _ := randomUser(t)
	pet := randomPet(t, user)

	testCases := []struct {
		name          string
		body          gin.H
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockAction)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"id":                pet.ID,
				"name":              pet.Name,
				"about":             pet.About.String,
				"user_id":           pet.UserID,
				"age":               pet.Age,
				"gender":            pet.Gender,
				"pet_type":          pet.PetType,
				"breed":             pet.Breed.String,
				"image_url":         pet.ImageUrl.String,
				"medical_condition": pet.MedicalCondition.String,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			},
			buildStubs: func(action *mockdb.MockAction) {
				arg := db.UpdatePetParams{
					Name:             pet.Name,
					About:            pet.About,
					Age:              pet.Age,
					Gender:           pet.Gender,
					PetType:          pet.PetType,
					Breed:            pet.Breed,
					ImageUrl:         pet.ImageUrl,
					MedicalCondition: pet.MedicalCondition,
					ID:               pet.ID,
				}
				action.EXPECT().
					UpdatePet(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(pet, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "BadRequest",
			body: gin.H{
				"id":                pet.ID,
				"name":              pet.Name,
				"about":             pet.About.String,
				"age":               pet.Age,
				"gender":            pet.Gender,
				"pet_type":          pet.PetType,
				"breed":             pet.Breed.String,
				"image_url":         pet.ImageUrl.String,
				"medical_condition": pet.MedicalCondition.String,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			},
			buildStubs: func(action *mockdb.MockAction) {
				arg := db.UpdatePetParams{
					Name:             pet.Name,
					About:            pet.About,
					Age:              pet.Age,
					Gender:           pet.Gender,
					PetType:          pet.PetType,
					Breed:            pet.Breed,
					ImageUrl:         pet.ImageUrl,
					MedicalCondition: pet.MedicalCondition,
					ID:               pet.ID,
				}
				action.EXPECT().
					UpdatePet(gomock.Any(), gomock.Eq(arg)).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		}, {
			name: "Unauthorized",
			body: gin.H{
				"id":                pet.ID,
				"name":              pet.Name,
				"about":             pet.About.String,
				"age":               pet.Age,
				"user_id":           pet.UserID + 1,
				"gender":            pet.Gender,
				"pet_type":          pet.PetType,
				"breed":             pet.Breed.String,
				"image_url":         pet.ImageUrl.String,
				"medical_condition": pet.MedicalCondition.String,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			},
			buildStubs: func(action *mockdb.MockAction) {
				arg := db.UpdatePetParams{
					Name:             pet.Name,
					About:            pet.About,
					Age:              pet.Age,
					Gender:           pet.Gender,
					PetType:          pet.PetType,
					Breed:            pet.Breed,
					ImageUrl:         pet.ImageUrl,
					MedicalCondition: pet.MedicalCondition,
					ID:               pet.ID,
				}
				action.EXPECT().
					UpdatePet(gomock.Any(), gomock.Eq(arg)).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockAction(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/users/pets/update"
			request, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}
func TestDeletePet(t *testing.T) {
	user, _ := randomUser(t)
	pet := randomPet(t, user)

	testCases := []struct {
		name          string
		body          gin.H
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockAction)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"id":      pet.ID,
				"user_id": pet.UserID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			},
			buildStubs: func(action *mockdb.MockAction) {
				action.EXPECT().
					DeletePet(gomock.Any(), gomock.Eq(pet.ID)).
					Times(1)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "BadRequest",
			body: gin.H{
				"id": pet.ID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			},
			buildStubs: func(action *mockdb.MockAction) {
				action.EXPECT().
					DeletePet(gomock.Any(), gomock.Eq(pet.ID)).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Unauthorized",
			body: gin.H{
				"id":      pet.ID,
				"user_id": pet.UserID + 1,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			},
			buildStubs: func(action *mockdb.MockAction) {
				action.EXPECT().
					DeletePet(gomock.Any(), gomock.Eq(pet.ID)).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"id":      pet.ID,
				"user_id": pet.UserID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			},
			buildStubs: func(action *mockdb.MockAction) {
				action.EXPECT().
					DeletePet(gomock.Any(), gomock.Any()).
					Times(1).
					Return(sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockAction(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/users/pets/delete"
			request, err := http.NewRequest(http.MethodDelete, url, bytes.NewReader(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func randomPet(t *testing.T, user db.User) (pet db.Pet) {
	pet = db.Pet{
		ID:               util.RandomInt(1, 1000),
		Name:             util.RandomString(10),
		About:            util.StringToNullString(util.RandomString(30)),
		UserID:           user.ID,
		Age:              int32(util.RandomAge()),
		Gender:           util.RandomGender(),
		PetType:          util.RandomAnimal(),
		Breed:            util.StringToNullString(util.RandomString(8)),
		ImageUrl:         util.StringToNullString(util.RandomImageUrl()),
		MedicalCondition: util.StringToNullString(util.RandomString(10)),
	}
	return
}
