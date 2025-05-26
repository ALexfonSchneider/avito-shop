package merch

import (
	"github.com/google/uuid"
	"time"
)

type Purchase struct {
	ID       string
	UserID   string
	MerchID  string
	Quantity int32
	Amount   int64
	BoughtAt time.Time
}

func (p *Purchase) Validate() error {
	if p.Amount < 0 {
		return MerchAmountMustBePositive
	}
	if p.Quantity <= 0 {
		return MerchQuantityMustBeGratenThenZero
	}

	return nil
}

func NewPurchase(merchID, userID string, quantity int32, amount int64) *Purchase {
	return &Purchase{
		ID:       uuid.NewString(),
		MerchID:  merchID,
		UserID:   userID,
		Quantity: quantity,
		Amount:   amount,
	}
}
