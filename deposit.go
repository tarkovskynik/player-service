package player

import "fmt"

type Deposit struct {
	UserID    uint64 `json:"user_id"`
	DepositID int    `json:"deposit_id"`
	Amount    int    `json:"amount"`
	Token     string `json:"token"`
}

type DepositStat struct {
	DepositID int `json:"deposit_id"`
	BalanceBefore int `json:"balance_before"`
	DepositAmount int `json:"deposit_amount"`
	BalanceAfter int `json:"balance_after"`
	Time string `json:"time"`
}

func NewErrorDepositIDHasBeenCreated(id int) error {
	return fmt.Errorf("Deposit Id #%d has been created", id)
}