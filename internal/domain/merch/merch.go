package merch

import (
	"github.com/google/uuid"
	"time"
)

type Merch struct {
	Id          string
	Name        string
	Description string
	Price       int64
	CreatedAt   time.Time
}

func NewMerch(Name string, Description string, Price int64, CreatedAt time.Time) *Merch {
	return &Merch{
		Id:          uuid.NewString(),
		Name:        Name,
		Description: Description,
		Price:       Price,
		CreatedAt:   CreatedAt,
	}
}

func (m *Merch) Validate() error {
	if m.Price <= 0 {
		return MerchPriceMustBeGreaterThenZero
	}

	return nil
}
