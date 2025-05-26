package merch

import "errors"

var (
	MerchPriceMustBeGratenThenZero = errors.New("merch price must be greater than zero")
)

var (
	MerchAmountMustBePositive         = errors.New("amount must be positive")
	MerchQuantityMustBeGratenThenZero = errors.New("quantity must be greater than zero")
)
