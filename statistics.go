package player

type Statistics struct {
	DepositCount int `json:"deposit_count"`
	DepositSum   int `json:"deposit_sum"`
	BetCount     int `json:"bet_count"`
	BetSum       int `json:"bet_sum"`
	WinCount     int `json:"win_count"`
	WinSum       int `json:"win_sum"`
}
