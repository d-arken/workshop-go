package user_test

import (
	"errors"
	"github.com/d-arken/workshop-go/5_orm/user"
	mock_user "github.com/d-arken/workshop-go/5_orm/user/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock_user.NewMockRepositoryInterface(ctrl)
	svc := user.NewService(mockRepo)

	createReq := user.CreateUserRequest{
		Name: "Testevaldo Pires",
		Age:  40,
	}

	t.Run("should create without errors", func(t *testing.T) {
		mockRepo.EXPECT().Create(&createReq).Times(1).Return(nil)

		err := svc.Create(&createReq)
		assert.Nil(t, err)
	})

	t.Run("should return error due to repository error", func(t *testing.T) {
		expectedErr := errors.New("repository error")
		mockRepo.EXPECT().Create(&createReq).Times(1).Return(expectedErr)

		err := svc.Create(&createReq)
		assert.Equal(t, expectedErr, err)
	})

}
