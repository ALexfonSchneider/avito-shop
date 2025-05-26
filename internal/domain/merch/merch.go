package merch

type Merch struct {
	Id          string
	Name        string
	Description string
	Price       int64
}

func NewMerch(Name string, Description string, Price int64) *Merch {
	return &Merch{
		Name:        Name,
		Description: Description,
		Price:       Price,
	}
}

func (m *Merch) Validate(merch Merch) error {
	if merch.Price <= 0 {
		return MerchPriceMustBeGratenThenZero
	}

	return nil
}
