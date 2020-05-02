// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/akoridze/Desktop/NewsFeeder/services/newsFeedRepo.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	models "github.com/AnaKoridze/NewsFeeder/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockNewsFeedRepository is a mock of NewsFeedRepository interface
type MockNewsFeedRepository struct {
	ctrl     *gomock.Controller
	recorder *MockNewsFeedRepositoryMockRecorder
}

// MockNewsFeedRepositoryMockRecorder is the mock recorder for MockNewsFeedRepository
type MockNewsFeedRepositoryMockRecorder struct {
	mock *MockNewsFeedRepository
}

// NewMockNewsFeedRepository creates a new mock instance
func NewMockNewsFeedRepository(ctrl *gomock.Controller) *MockNewsFeedRepository {
	mock := &MockNewsFeedRepository{ctrl: ctrl}
	mock.recorder = &MockNewsFeedRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNewsFeedRepository) EXPECT() *MockNewsFeedRepositoryMockRecorder {
	return m.recorder
}

// GetAllFeeds mocks base method
func (m *MockNewsFeedRepository) GetAllFeeds() (*models.GetNewsFeedsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFeeds")
	ret0, _ := ret[0].(*models.GetNewsFeedsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllFeeds indicates an expected call of GetAllFeeds
func (mr *MockNewsFeedRepositoryMockRecorder) GetAllFeeds() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFeeds", reflect.TypeOf((*MockNewsFeedRepository)(nil).GetAllFeeds))
}

// CreateFeed mocks base method
func (m *MockNewsFeedRepository) CreateFeed(request models.CreateNewsFeedRequest) (*models.CreateNewsFeedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFeed", request)
	ret0, _ := ret[0].(*models.CreateNewsFeedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFeed indicates an expected call of CreateFeed
func (mr *MockNewsFeedRepositoryMockRecorder) CreateFeed(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFeed", reflect.TypeOf((*MockNewsFeedRepository)(nil).CreateFeed), request)
}