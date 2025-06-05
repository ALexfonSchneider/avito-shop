package dto

type SendCoinRequest struct {
	FromUserID string
	ToUserName string
	Amount     int64
}
