package user

func (u *User) CanBuy(amount int64) bool {
	return u.Balance >= amount
}

func (u *User) CanSend(amount int64) bool {
	return u.Balance >= amount
}
