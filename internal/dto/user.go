package dto

type UserInfo struct {
	Coins     int64
	Inventory []InventoryItem
	History   CoinHistory
}

type InventoryItem struct {
	Name     string
	Quantity int32
}

type CoinHistory struct {
	Received []CoinReceived
	Sent     []CoinSent
}

type CoinReceived struct {
	FromUser string
	Amount   int64
}

type CoinSent struct {
	ToUser string
	Amount int64
}
