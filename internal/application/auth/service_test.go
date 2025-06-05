package auth_test

import (
	"github.com/ALexfonSchneider/avito-shop/internal/application/auth"
	userdomain "github.com/ALexfonSchneider/avito-shop/internal/domain/user"
	"github.com/ALexfonSchneider/avito-shop/internal/dto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"os"
	"testing"
	"time"
)

var user1 = userdomain.User{
	ID:        "0",
	Username:  "TestUser1",
	Password:  "password",
	Balance:   1000,
	CreatedAt: time.Now(),
}

func prepare(t *testing.T) (*gomock.Controller, *MockUserRepository, *MockJWTProvider, *MockHasher, *auth.Service) {
	ctrl := gomock.NewController(t)
	mockedUserRepo := NewMockUserRepository(ctrl)
	mockedJWTProvider := NewMockJWTProvider(ctrl)
	mockedHasher := NewMockHasher(ctrl)

	l := slog.New(slog.NewTextHandler(os.Stdout, nil))

	mockedService := auth.New(mockedUserRepo, mockedJWTProvider, mockedHasher, l)

	return ctrl, mockedUserRepo, mockedJWTProvider, mockedHasher, mockedService
}

func TestService_Authorize_UserNotExists_UserCreate_Success(t *testing.T) {
	ctrl, mockedUserRepo, mockedJWTProvider, mockedHasher, mockedService := prepare(t)
	defer ctrl.Finish()

	request := dto.AuthorizeRequest{
		Username: user1.Username,
		Password: user1.Password,
	}

	mockedUserRepo.EXPECT().FindUserByUsername(gomock.Any(), request.Username).Return(nil, nil)
	mockedHasher.EXPECT().Hash(request.Password).Return(request.Password, nil)
	mockedUserRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(nil)
	mockedHasher.EXPECT().Compare(request.Password, user1.Password).Return(nil)
	mockedJWTProvider.EXPECT().CreateToken(gomock.Any()).Return("token", nil)

	token, err := mockedService.Authorize(t.Context(), request)

	assert.NoError(t, err)
	assert.Equal(t, token, "token")
}
