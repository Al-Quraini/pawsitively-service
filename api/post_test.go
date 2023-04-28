package api

import (
	"bytes"
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

func TestCreatePost(t *testing.T) {
	user, _ := randomUser(t)
	post := randomPost(t, user)

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
				"title":     post.Title.String,
				"body":      post.Body.String,
				"user_id":   post.UserID,
				"image_url": post.ImageUrl.String,
				"status":    post.Status.String,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			},
			buildStubs: func(action *mockdb.MockAction) {
				arg := db.CreatePostParams{
					Title:    post.Title,
					Body:     post.Body,
					UserID:   post.UserID,
					ImageUrl: post.ImageUrl,
					Status:   post.Status,
				}
				action.EXPECT().
					CreatePost(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(post, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "BadRequest",
			body: gin.H{
				"title":     1,
				"body":      post.Body.String,
				"user_id":   post.UserID,
				"image_url": post.ImageUrl.String,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			},
			buildStubs: func(action *mockdb.MockAction) {
				arg := db.CreatePostParams{
					Title:    post.Title,
					Body:     post.Body,
					UserID:   post.UserID,
					ImageUrl: post.ImageUrl,
					Status:   post.Status,
				}
				action.EXPECT().
					CreatePost(gomock.Any(), gomock.Eq(arg)).
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

			url := "/posts/create"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func TestGetPost(t *testing.T) {
	user, _ := randomUser(t)
	post := randomPost(t, user)

	testCases := []struct {
		name          string
		ID            int64
		buildStubs    func(action *mockdb.MockAction)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			ID:   post.ID,
			buildStubs: func(action *mockdb.MockAction) {
				action.EXPECT().
					GetPost(gomock.Any(), gomock.Any()).
					Times(1).
					Return(post, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "BadRequest",
			ID:   -1,
			buildStubs: func(action *mockdb.MockAction) {
				action.EXPECT().
					GetPost(gomock.Any(), gomock.Any()).
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

			url := fmt.Sprintf("/posts/%d", tc.ID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func TestListPosts(t *testing.T) {
	user, _ := randomUser(t)
	// var lastPost db.Post

	n := 5
	posts := make([]db.Post, n)
	for i := 0; i < n; i++ {
		posts[i] = randomPost(t, user)
		// lastPost = posts[i]
	}

	type Query struct {
		offset int
		limit  int
	}

	testCases := []struct {
		name          string
		query         Query
		buildStubs    func(action *mockdb.MockAction)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			query: Query{
				offset: 0,
				limit:  -1,
			},
			buildStubs: func(action *mockdb.MockAction) {

				arg := db.ListPostsParams{
					Limit:  -1,
					Offset: 0,
				}
				action.EXPECT().
					ListPosts(gomock.Any(), arg).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "OK",
			query: Query{
				offset: -1,
				limit:  n,
			},
			buildStubs: func(action *mockdb.MockAction) {

				arg := db.ListPostsParams{
					Limit:  int32(n),
					Offset: -1,
				}
				action.EXPECT().
					ListPosts(gomock.Any(), arg).
					Times(1).
					Return(posts, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
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

			url := "/posts"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			q := request.URL.Query()
			q.Add("limit", fmt.Sprintf("%d", tc.query.limit))
			q.Add("offset", fmt.Sprintf("%d", tc.query.offset))
			request.URL.RawQuery = q.Encode()

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func TestListPostsByUserID(t *testing.T) {
	user, _ := randomUser(t)
	var lastPost db.Post

	n := 5
	posts := make([]db.Post, n)
	for i := 0; i < n; i++ {
		posts[i] = randomPost(t, user)
		lastPost = posts[i]
	}

	type Query struct {
		ID     int
		Offset int
		Limit  int
	}

	testCases := []struct {
		name          string
		query         Query
		buildStubs    func(action *mockdb.MockAction)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			query: Query{
				ID:     int(lastPost.UserID),
				Offset: 0,
				Limit:  n,
			},
			buildStubs: func(action *mockdb.MockAction) {

				arg := db.ListPostsByUserIDParams{
					UserID: lastPost.UserID,
					Limit:  int32(n),
					Offset: 0,
				}
				action.EXPECT().
					ListPostsByUserID(gomock.Any(), arg).
					Times(1).
					Return(posts, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "BadRequest1",
			query: Query{
				ID:     int(lastPost.UserID),
				Offset: 0,
				Limit:  -n,
			},
			buildStubs: func(action *mockdb.MockAction) {

				arg := db.ListPostsByUserIDParams{
					UserID: lastPost.UserID,
					Limit:  int32(-n),
					Offset: 0,
				}
				action.EXPECT().
					ListPostsByUserID(gomock.Any(), arg).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "BadRequest2",
			query: Query{
				ID:     -1,
				Offset: 0,
				Limit:  n,
			},
			buildStubs: func(action *mockdb.MockAction) {

				arg := db.ListPostsByUserIDParams{
					UserID: lastPost.UserID,
					Limit:  int32(n),
					Offset: 0,
				}
				action.EXPECT().
					ListPostsByUserID(gomock.Any(), arg).
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

			url := fmt.Sprintf("/users/%d/posts", tc.query.ID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			q := request.URL.Query()
			q.Add("limit", fmt.Sprintf("%d", tc.query.Limit))
			q.Add("offset", fmt.Sprintf("%d", tc.query.Offset))
			request.URL.RawQuery = q.Encode()

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func TestUpdatePost(t *testing.T) {
	user, _ := randomUser(t)
	post := randomPost(t, user)

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
				"id":        post.ID,
				"title":     post.Title.String,
				"body":      post.Body.String,
				"user_id":   post.UserID,
				"image_url": post.ImageUrl.String,
				"status":    post.Status.String,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			},
			buildStubs: func(action *mockdb.MockAction) {
				arg := db.UpdatePostParams{
					Title:    post.Title,
					Body:     post.Body,
					ImageUrl: post.ImageUrl,
					Status:   post.Status,
					ID:       post.ID,
				}
				action.EXPECT().
					UpdatePost(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(post, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "BadRequest",
			body: gin.H{
				"id":        "1",
				"title":     post.Title.String,
				"body":      post.Body.String,
				"user_id":   post.UserID,
				"image_url": post.ImageUrl.String,
				"status":    post.Status.String,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			},
			buildStubs: func(action *mockdb.MockAction) {
				arg := db.UpdatePostParams{
					Title:    post.Title,
					Body:     post.Body,
					ImageUrl: post.ImageUrl,
					Status:   post.Status,
					ID:       post.ID,
				}
				action.EXPECT().
					UpdatePost(gomock.Any(), gomock.Eq(arg)).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Unauthorized",
			body: gin.H{
				"id":        post.ID,
				"title":     post.Title.String,
				"body":      post.Body.String,
				"user_id":   post.UserID + 1,
				"image_url": post.ImageUrl.String,
				"status":    post.Status.String,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			},
			buildStubs: func(action *mockdb.MockAction) {
				arg := db.UpdatePostParams{
					Title:    post.Title,
					Body:     post.Body,
					ImageUrl: post.ImageUrl,
					Status:   post.Status,
					ID:       post.ID,
				}
				action.EXPECT().
					UpdatePost(gomock.Any(), gomock.Eq(arg)).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		// {
		// 	name: "BadRequest",
		// 	body: gin.H{
		// 		"title":     1,
		// 		"body":      post.Body.String,
		// 		"user_id":   post.UserID,
		// 		"image_url": post.ImageUrl.String,
		// 	},
		// 	setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
		// 		addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
		// 	},
		// 	buildStubs: func(action *mockdb.MockAction) {
		// 		arg := db.CreatePostParams{
		// 			Title:    post.Title,
		// 			Body:     post.Body,
		// 			UserID:   post.UserID,
		// 			ImageUrl: post.ImageUrl,
		// 			Status:   post.Status,
		// 		}
		// 		action.EXPECT().
		// 			CreatePost(gomock.Any(), gomock.Eq(arg)).
		// 			Times(0)
		// 	},
		// 	checkResponse: func(recorder *httptest.ResponseRecorder) {
		// 		require.Equal(t, http.StatusBadRequest, recorder.Code)
		// 	},
		// },
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

			url := "/posts/update"
			request, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func TestDeletePost(t *testing.T) {
	user, _ := randomUser(t)
	post := randomPost(t, user)

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
				"id":      post.ID,
				"user_id": post.UserID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			},
			buildStubs: func(action *mockdb.MockAction) {
				action.EXPECT().
					DeletePost(gomock.Any(), gomock.Eq(post.ID)).
					Times(1)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "BadRequest",
			body: gin.H{
				"id":      "-1",
				"user_id": post.UserID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			},
			buildStubs: func(action *mockdb.MockAction) {
				action.EXPECT().
					DeletePost(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Unauthorized",
			body: gin.H{
				"id":      post.ID,
				"user_id": post.UserID + 1,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, time.Minute)
			},
			buildStubs: func(action *mockdb.MockAction) {
				action.EXPECT().
					DeletePost(gomock.Any(), gomock.Eq(post.ID)).
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

			action := mockdb.NewMockAction(ctrl)
			tc.buildStubs(action)

			server := newTestServer(t, action)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/posts/delete"
			request, err := http.NewRequest(http.MethodDelete, url, bytes.NewReader(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}
func randomPost(t *testing.T, user db.User) (post db.Post) {
	post = db.Post{
		ID:         util.RandomInt(0, 20),
		Title:      util.StringToNullString(util.RandomString(10)),
		Body:       util.StringToNullString(util.RandomString(20)),
		UserID:     user.ID,
		ImageUrl:   util.StringToNullString(util.RandomImageUrl()),
		Status:     util.StringToNullString(util.RandomString(5)),
		LikesCount: 0,
	}
	return
}
