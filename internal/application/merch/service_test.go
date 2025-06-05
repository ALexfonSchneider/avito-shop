package merch_test

import (
	"github.com/ALexfonSchneider/avito-shop/internal/application/merch"
	domainmerch "github.com/ALexfonSchneider/avito-shop/internal/domain/merch"
	domainuser "github.com/ALexfonSchneider/avito-shop/internal/domain/user"
	"github.com/ALexfonSchneider/avito-shop/internal/dto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func prepare(t *testing.T) (*gomock.Controller, *MockUserRepository, *MockMerchRepository, *merch.Service) {
	ctrl := gomock.NewController(t)
	mockedUserRepository := NewMockUserRepository(ctrl)
	mockedMerchRepository := NewMockMerchRepository(ctrl)

	service := merch.New(mockedUserRepository, mockedMerchRepository)

	return ctrl, mockedUserRepository, mockedMerchRepository, service
}

func TestService_BuyMerch_Success(t *testing.T) {
	ctrl, mockedUserRepository, mockedMerchRepository, service := prepare(t)
	defer ctrl.Finish()

	buyingMerch := domainmerch.NewMerch("t-shirt", "", 100, time.Now())
	if err := buyingMerch.Validate(); err != nil {
		panic(err)
	}

	user := domainuser.NewUser("TestUser", "password", 1000, time.Now())
	if err := user.Validate(); err != nil {
		panic(err)
	}

	request := dto.BuyMerchRequest{
		MerchName: buyingMerch.Name,
		UserID:    user.ID,
		Quantity:  1,
	}

	purchase := domainmerch.NewPurchaseFromMerchAndUser(user.ID, *buyingMerch, request.Quantity, time.Now())
	if err := purchase.Validate(); err != nil {
		panic(err)
	}

	mockedMerchRepository.EXPECT().FindMerchByName(gomock.Any(), request.MerchName).Return(buyingMerch, nil)
	mockedUserRepository.EXPECT().FindUserByID(gomock.Any(), user.ID).Return(user, nil)
	mockedUserRepository.EXPECT().IncrementUserBalance(
		gomock.Any(), user.ID, -purchase.Amount,
	).Return(nil)
	mockedMerchRepository.EXPECT().CreatePurchase(gomock.Any(), gomock.Any()).Return(nil)

	err := service.BuyMerch(t.Context(), request)

	assert.NoError(t, err)
}
