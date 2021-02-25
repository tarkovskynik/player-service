package player

type Deposit struct {
	UserID uint64 `json:"user_id"`
	DepositID int `json:"deposit_id"`
	Amount int `json:"amount"`
	Token string `json:"token"`
}
