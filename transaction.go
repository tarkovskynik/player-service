package player

import "fmt"

type Transaction struct {
	UserID        uint64 `json:"user_id"`
	TransactionId int    `json:"transaction_id"`
	Type          string `json:"type"`
	Amount        int    `json:"amount"`
	Token         string `json:"token"`
}

type TransactionStat struct {
	TransactionId int `json:"transaction_id"`
	BalanceBefore int `json:"balance_before"`
	TransactionAmount int `json:"transaction_amount"`
	BalanceAfter int `json:"balance_after"`
	Time string `json:"time"`
}

func NewErrorTransactionIDHasBeenCreated(id int) error {
	return fmt.Errorf("Transaction Id #%d has been created", id)
}