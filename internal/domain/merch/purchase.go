package merch

import (
	"github.com/google/uuid"
	"time"
)

type Purchase struct {
	ID          string
	UserID      string
	MerchID     string
	Quantity    int32
	Amount      int64
	PurchasedAt time.Time
}

func (p *Purchase) Validate() error {
	if p.Amount < 0 {
		return PurchaseAmountMustBePositive
	}
	if p.Quantity <= 0 {
		return PurchaseQuantityMustBeGreaterThenZero
	}

	return nil
}

func NewPurchase(merchID, userID string, quantity int32, amount int64, purchasedAt time.Time) *Purchase {
	return &Purchase{
		ID:          uuid.NewString(),
		MerchID:     merchID,
		UserID:      userID,
		Quantity:    quantity,
		Amount:      amount,
		PurchasedAt: purchasedAt,
	}
}

func NewPurchaseFromMerchAndUser(userID string, merch Merch, quantity int32, purchasedAt time.Time) *Purchase {
	return &Purchase{
		ID:          uuid.NewString(),
		UserID:      userID,
		MerchID:     merch.Id,
		Quantity:    quantity,
		Amount:      merch.Price * int64(quantity),
		PurchasedAt: purchasedAt,
	}
}
