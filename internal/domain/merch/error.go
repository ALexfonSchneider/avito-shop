package merch

import "errors"

var (
	MerchPriceMustBeGreaterThenZero = errors.New("merch price must be greater than zero")
)

var (
	PurchaseAmountMustBePositive          = errors.New("amount must be positive")
	PurchaseQuantityMustBeGreaterThenZero = errors.New("quantity must be greater than zero")
)
