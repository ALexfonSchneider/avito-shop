package user_test

import (
	"github.com/ALexfonSchneider/avito-shop/internal/application/user"
	transactiondomain "github.com/ALexfonSchneider/avito-shop/internal/domain/transaction"
	userdomain "github.com/ALexfonSchneider/avito-shop/internal/domain/user"
	"github.com/ALexfonSchneider/avito-shop/internal/dto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func prepare(t *testing.T) (*gomock.Controller, *MockTransactionRepository, *MockUserRepository, *user.Service) {
	ctrl := gomock.NewController(t)

	mockedTransactionRepository := NewMockTransactionRepository(ctrl)
	mockedUserRepository := NewMockUserRepository(ctrl)

	service := user.New(mockedUserRepository, mockedTransactionRepository)

	return ctrl, mockedTransactionRepository, mockedUserRepository, service
}

func TestService_GetUserInfo(t *testing.T) {
	ctrl, mockedTransactionRepository, mockedUserRepository, service := prepare(t)
	defer ctrl.Finish()

	currentUser := userdomain.NewUser("TestUser1", "password", 1000, time.Now())
	if err := currentUser.Validate(); err != nil {
		panic(err)
	}
	otherUser := userdomain.NewUser("TestUser2", "password", 1000, time.Now())
	if err := otherUser.Validate(); err != nil {
		panic(err)
	}

	mockedUserRepository.EXPECT().FindUserByID(gomock.Any(), currentUser.ID).Return(currentUser, nil)

	inventory := []dto.InventoryItem{
		{"merch1", 1},
		{"merch2", 2},
		{"merch3", 3},
	}

	mockedUserRepository.EXPECT().UserInventory(gomock.Any(), currentUser.ID).Return(inventory, nil)

	transactions := []transactiondomain.Transaction{
		*transactiondomain.NewTransaction(currentUser.ID, otherUser.ID, 100),
		*transactiondomain.NewTransaction(otherUser.ID, currentUser.ID, 50),
		*transactiondomain.NewTransaction(currentUser.ID, otherUser.ID, 100),
		*transactiondomain.NewTransaction(currentUser.ID, otherUser.ID, 300),
	}

	for _, transaction := range transactions {
		if err := transaction.Validate(); err != nil {
			panic(err)
		}
	}

	mockedTransactionRepository.EXPECT().GetTransactions(gomock.Any(), currentUser.ID).Return(transactions, nil)

	info, err := service.GetUserInfo(t.Context(), currentUser.ID)

	assert.NoError(t, err)

	assert.Equal(t, len(info.Inventory), len(inventory))
	assert.Equal(t, len(info.History.Received), 1)
	assert.Equal(t, len(info.History.Sent), 3)
}
