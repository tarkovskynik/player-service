package player

type Transaction struct {
	UserID        uint64 `json:"user_id"`
	TransactionId int    `json:"transaction_id"`
	Type          string `json:"type"`
	Amount        int    `json:"amount"`
	Token         string `json:"token"`
}

type TransactionStat struct {
	TransactionId int `json:"transaction_id"`
	TransactionAmount int `json:"transaction_amount"`
	BalanceBefore int `json:"balance_before"`
	BalanceAfter int `json:"balance_after"`
}
