package transaction

import (
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestTransaction_Validate(t1 *testing.T) {
	id := uuid.NewString()

	tests := []struct {
		name        string
		transaction Transaction
		wantErr     bool
	}{
		{
			name: "valid object",
			transaction: Transaction{
				SenderID:   uuid.NewString(),
				ReceiverID: uuid.NewString(),
				Amount:     100,
				CreatedAt:  time.Now(),
			},
			wantErr: false,
		},
		{
			name: "invalid object. Amount is less than zero",
			transaction: Transaction{
				SenderID:   uuid.NewString(),
				ReceiverID: uuid.NewString(),
				Amount:     -10,
				CreatedAt:  time.Now(),
			},
			wantErr: true,
		},
		{
			name: "invalid object. Cannot send to itself",
			transaction: Transaction{
				SenderID:   id,
				ReceiverID: id,
				Amount:     100,
				CreatedAt:  time.Now(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Transaction{
				SenderID:   tt.transaction.SenderID,
				ReceiverID: tt.transaction.ReceiverID,
				Amount:     tt.transaction.Amount,
				CreatedAt:  tt.transaction.CreatedAt,
			}
			if err := t.Validate(); (err != nil) != tt.wantErr {
				t1.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
