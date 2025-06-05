package transaction_test

import (
	"github.com/ALexfonSchneider/avito-shop/internal/application/transaction"
	userdomain "github.com/ALexfonSchneider/avito-shop/internal/domain/user"
	"github.com/ALexfonSchneider/avito-shop/internal/dto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func prepare(t *testing.T) (*gomock.Controller, *MockUserRepository, *MockTransactionRepository, *transaction.Service) {
	ctrl := gomock.NewController(t)

	mockUserRepository := NewMockUserRepository(ctrl)
	mockTransactionRepository := NewMockTransactionRepository(ctrl)

	service := transaction.New(mockTransactionRepository, mockUserRepository)

	return ctrl, mockUserRepository, mockTransactionRepository, service
}

func TestService_SendCoin_Success(t *testing.T) {
	ctrl, mockUserRepository, mockTransactionRepository, service := prepare(t)
	defer ctrl.Finish()

	fromUser := userdomain.NewUser("TestUser1", "passsword", 1000, time.Now())
	if err := fromUser.Validate(); err != nil {
		panic(err)
	}

	toUser := userdomain.NewUser("TestUser2", "passsword", 1000, time.Now())
	if err := toUser.Validate(); err != nil {
		panic(err)
	}

	request := dto.SendCoinRequest{
		FromUserID: fromUser.ID,
		ToUserName: toUser.Username,
		Amount:     100,
	}

	mockUserRepository.EXPECT().FindUserByID(gomock.Any(), fromUser.ID).Return(fromUser, nil)
	mockUserRepository.EXPECT().FindUserByUsername(gomock.Any(), toUser.Username).Return(toUser, nil)

	mockUserRepository.EXPECT().IncrementUserBalance(t.Context(), fromUser.ID, -request.Amount).Return(nil)
	mockUserRepository.EXPECT().IncrementUserBalance(t.Context(), toUser.ID, request.Amount).Return(nil)

	mockTransactionRepository.EXPECT().CreateTransaction(gomock.Any(), gomock.Any()).Return(nil)

	err := service.SendCoin(t.Context(), request)

	assert.NoError(t, err)
}
