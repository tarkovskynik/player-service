package player

type Transaction struct{
	UserID uint64 `json:"user_id"`
	TransactionId int `json:"transaction_id"`
	Type string `json:"type"`
	Amount int `json:"amount"`
	Token string `json:"token"`
}
