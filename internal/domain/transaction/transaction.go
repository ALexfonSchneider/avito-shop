package transaction

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	ID         string
	SenderID   string
	ReceiverID string
	Amount     int64
	CreatedAt  time.Time
}

func (t *Transaction) Validate() error {
	if t.SenderID == t.ReceiverID {
		return CannotSendCoinsToYourSelf
	}
	if t.Amount <= 0 {
		return AmountMustBeGreaterThanZero
	}
	return nil
}

func NewTransaction(senderID string, receiverID string, amount int64) *Transaction {
	return &Transaction{
		ID:         uuid.NewString(),
		SenderID:   senderID,
		ReceiverID: receiverID,
		Amount:     amount,
	}
}
