package handlers

import (
	"github.com/AnaKoridze/NewsFeeder/mocks"
	"github.com/AnaKoridze/NewsFeeder/models"
	"github.com/AnaKoridze/NewsFeeder/services"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestPostNewsFeedSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	newsFeedServiceMock := mock_services.NewMockNewsFeedRepository(ctrl)

	// When
	requestBody := models.CreateNewsFeedRequest{
		Title:  "Ana",
		Post:   "Bana",
		Author: "Me",
	}
	expected := models.CreateNewsFeedResponse{ID: 1}
	newsFeedServiceMock.EXPECT().CreateFeed(gomock.Eq(requestBody)).Return(&expected, nil)

	// Verify
	s := services.Services{NewsFeedService: newsFeedServiceMock}
	resp, err := s.NewsFeedService.CreateFeed(requestBody)

	if err != nil {
		t.Fatal("Create Feed failed", err)
	}

	if !reflect.DeepEqual(resp, &expected) {
		t.Fatalf("Wrong Response expected %v found %v", expected, resp)
	}
}
