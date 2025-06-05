package merch

import (
	"testing"
	"time"
)

func TestMerch_Validate(t *testing.T) {
	tests := []struct {
		name    string
		merch   Merch
		wantErr bool
	}{
		{
			name: "valid object",
			merch: Merch{
				Name:        "t-shirt",
				Description: "cool t-shirt",
				Price:       1000,
			},
			wantErr: false,
		},
		{
			name: "invalid object. Price is less than zero",
			merch: Merch{
				Name:        "t-shirt",
				Description: "cool t-shirt",
				Price:       -100,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Merch{
				Name:        tt.merch.Name,
				Description: tt.merch.Description,
				Price:       tt.merch.Price,
			}
			if err := m.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPurchase_Validate(t *testing.T) {
	tests := []struct {
		name     string
		purchase Purchase
		wantErr  bool
	}{
		{
			name: "valid object",
			purchase: Purchase{
				Quantity:    100,
				Amount:      100 * 1000,
				PurchasedAt: time.Now(),
			},
			wantErr: false,
		},
		{
			name: "invalid object. quantity less then zero",
			purchase: Purchase{
				Quantity:    -1,
				Amount:      100 * 1000,
				PurchasedAt: time.Now(),
			},
			wantErr: true,
		},
		{
			name: "invalid object. amount less then zero",
			purchase: Purchase{
				Quantity:    100,
				Amount:      -100 * 1000,
				PurchasedAt: time.Now(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Purchase{
				Quantity:    tt.purchase.Quantity,
				Amount:      tt.purchase.Amount,
				PurchasedAt: tt.purchase.PurchasedAt,
			}
			if err := p.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
