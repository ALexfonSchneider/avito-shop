package merch

import (
	"errors"
	"github.com/ALexfonSchneider/avito-shop/internal/domain"
)

var (
	MerchPriceMustBeGreaterThenZero = domain.Error{Err: errors.New("merch price must be greater than zero")}
)

var (
	PurchaseAmountMustBePositive          = domain.Error{Err: errors.New("amount must be positive")}
	PurchaseQuantityMustBeGreaterThenZero = domain.Error{Err: errors.New("purchase quantity must be greater than zero")}
)
