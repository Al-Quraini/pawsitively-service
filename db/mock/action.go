// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/alquraini/pawsitively/db/sqlc (interfaces: Action)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	db "github.com/alquraini/pawsitively/db/sqlc"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockAction is a mock of Action interface.
type MockAction struct {
	ctrl     *gomock.Controller
	recorder *MockActionMockRecorder
}

// MockActionMockRecorder is the mock recorder for MockAction.
type MockActionMockRecorder struct {
	mock *MockAction
}

// NewMockAction creates a new mock instance.
func NewMockAction(ctrl *gomock.Controller) *MockAction {
	mock := &MockAction{ctrl: ctrl}
	mock.recorder = &MockActionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAction) EXPECT() *MockActionMockRecorder {
	return m.recorder
}

// CreateLike mocks base method.
func (m *MockAction) CreateLike(arg0 context.Context, arg1 db.CreateLikeParams) (db.Like, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLike", arg0, arg1)
	ret0, _ := ret[0].(db.Like)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLike indicates an expected call of CreateLike.
func (mr *MockActionMockRecorder) CreateLike(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLike", reflect.TypeOf((*MockAction)(nil).CreateLike), arg0, arg1)
}

// CreatePet mocks base method.
func (m *MockAction) CreatePet(arg0 context.Context, arg1 db.CreatePetParams) (db.Pet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePet", arg0, arg1)
	ret0, _ := ret[0].(db.Pet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePet indicates an expected call of CreatePet.
func (mr *MockActionMockRecorder) CreatePet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePet", reflect.TypeOf((*MockAction)(nil).CreatePet), arg0, arg1)
}

// CreatePost mocks base method.
func (m *MockAction) CreatePost(arg0 context.Context, arg1 db.CreatePostParams) (db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", arg0, arg1)
	ret0, _ := ret[0].(db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePost indicates an expected call of CreatePost.
func (mr *MockActionMockRecorder) CreatePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockAction)(nil).CreatePost), arg0, arg1)
}

// CreateSession mocks base method.
func (m *MockAction) CreateSession(arg0 context.Context, arg1 db.CreateSessionParams) (db.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", arg0, arg1)
	ret0, _ := ret[0].(db.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockActionMockRecorder) CreateSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockAction)(nil).CreateSession), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockAction) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockActionMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAction)(nil).CreateUser), arg0, arg1)
}

// DeleteLike mocks base method.
func (m *MockAction) DeleteLike(arg0 context.Context, arg1 db.DeleteLikeParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLike", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLike indicates an expected call of DeleteLike.
func (mr *MockActionMockRecorder) DeleteLike(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLike", reflect.TypeOf((*MockAction)(nil).DeleteLike), arg0, arg1)
}

// DeletePet mocks base method.
func (m *MockAction) DeletePet(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePet", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePet indicates an expected call of DeletePet.
func (mr *MockActionMockRecorder) DeletePet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePet", reflect.TypeOf((*MockAction)(nil).DeletePet), arg0, arg1)
}

// DeletePost mocks base method.
func (m *MockAction) DeletePost(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePost", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePost indicates an expected call of DeletePost.
func (mr *MockActionMockRecorder) DeletePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*MockAction)(nil).DeletePost), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockAction) DeleteUser(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockActionMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockAction)(nil).DeleteUser), arg0, arg1)
}

// GetLikeFromPostForUser mocks base method.
func (m *MockAction) GetLikeFromPostForUser(arg0 context.Context, arg1 db.GetLikeFromPostForUserParams) (db.Like, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLikeFromPostForUser", arg0, arg1)
	ret0, _ := ret[0].(db.Like)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLikeFromPostForUser indicates an expected call of GetLikeFromPostForUser.
func (mr *MockActionMockRecorder) GetLikeFromPostForUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLikeFromPostForUser", reflect.TypeOf((*MockAction)(nil).GetLikeFromPostForUser), arg0, arg1)
}

// GetLikesFromPost mocks base method.
func (m *MockAction) GetLikesFromPost(arg0 context.Context, arg1 int64) ([]db.Like, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLikesFromPost", arg0, arg1)
	ret0, _ := ret[0].([]db.Like)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLikesFromPost indicates an expected call of GetLikesFromPost.
func (mr *MockActionMockRecorder) GetLikesFromPost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLikesFromPost", reflect.TypeOf((*MockAction)(nil).GetLikesFromPost), arg0, arg1)
}

// GetPetById mocks base method.
func (m *MockAction) GetPetById(arg0 context.Context, arg1 int64) (db.Pet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPetById", arg0, arg1)
	ret0, _ := ret[0].(db.Pet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPetById indicates an expected call of GetPetById.
func (mr *MockActionMockRecorder) GetPetById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPetById", reflect.TypeOf((*MockAction)(nil).GetPetById), arg0, arg1)
}

// GetPets mocks base method.
func (m *MockAction) GetPets(arg0 context.Context, arg1 int64) ([]db.Pet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPets", arg0, arg1)
	ret0, _ := ret[0].([]db.Pet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPets indicates an expected call of GetPets.
func (mr *MockActionMockRecorder) GetPets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPets", reflect.TypeOf((*MockAction)(nil).GetPets), arg0, arg1)
}

// GetPost mocks base method.
func (m *MockAction) GetPost(arg0 context.Context, arg1 int64) (db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPost", arg0, arg1)
	ret0, _ := ret[0].(db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPost indicates an expected call of GetPost.
func (mr *MockActionMockRecorder) GetPost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPost", reflect.TypeOf((*MockAction)(nil).GetPost), arg0, arg1)
}

// GetSession mocks base method.
func (m *MockAction) GetSession(arg0 context.Context, arg1 uuid.UUID) (db.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", arg0, arg1)
	ret0, _ := ret[0].(db.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockActionMockRecorder) GetSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockAction)(nil).GetSession), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockAction) GetUser(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockActionMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockAction)(nil).GetUser), arg0, arg1)
}

// GetUserByID mocks base method.
func (m *MockAction) GetUserByID(arg0 context.Context, arg1 int64) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockActionMockRecorder) GetUserByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockAction)(nil).GetUserByID), arg0, arg1)
}

// ListPosts mocks base method.
func (m *MockAction) ListPosts(arg0 context.Context, arg1 db.ListPostsParams) ([]db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPosts", arg0, arg1)
	ret0, _ := ret[0].([]db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPosts indicates an expected call of ListPosts.
func (mr *MockActionMockRecorder) ListPosts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPosts", reflect.TypeOf((*MockAction)(nil).ListPosts), arg0, arg1)
}

// ListPostsByUserID mocks base method.
func (m *MockAction) ListPostsByUserID(arg0 context.Context, arg1 db.ListPostsByUserIDParams) ([]db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPostsByUserID", arg0, arg1)
	ret0, _ := ret[0].([]db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPostsByUserID indicates an expected call of ListPostsByUserID.
func (mr *MockActionMockRecorder) ListPostsByUserID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPostsByUserID", reflect.TypeOf((*MockAction)(nil).ListPostsByUserID), arg0, arg1)
}

// UpdatePet mocks base method.
func (m *MockAction) UpdatePet(arg0 context.Context, arg1 db.UpdatePetParams) (db.Pet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePet", arg0, arg1)
	ret0, _ := ret[0].(db.Pet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePet indicates an expected call of UpdatePet.
func (mr *MockActionMockRecorder) UpdatePet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePet", reflect.TypeOf((*MockAction)(nil).UpdatePet), arg0, arg1)
}

// UpdatePost mocks base method.
func (m *MockAction) UpdatePost(arg0 context.Context, arg1 db.UpdatePostParams) (db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePost", arg0, arg1)
	ret0, _ := ret[0].(db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePost indicates an expected call of UpdatePost.
func (mr *MockActionMockRecorder) UpdatePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*MockAction)(nil).UpdatePost), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockAction) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockActionMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockAction)(nil).UpdateUser), arg0, arg1)
}
