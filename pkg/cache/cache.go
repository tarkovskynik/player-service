package cache

import (
	"player"
	"sync"
	"time"
)

type Cache struct {
	data  map[uint64]*player.User
	datas map[uint64]*player.Statistics
	datad map[int]*player.DepositStat
	datat map[int]*player.TransactionStat
	sync.Mutex
}

func NewCache() *Cache {
	return &Cache{
		data:  make(map[uint64]*player.User),
		datas: make(map[uint64]*player.Statistics),
		datad: make(map[int]*player.DepositStat),
		datat: make(map[int]*player.TransactionStat),
	}
}

func (c *Cache) Create(user *player.User, statistics *player.Statistics) error {
	_, ok := c.data[user.Id]
	if ok {
		return player.NewErrorUserHasBeenCreated(user.Id)
	}
	c.Lock()
	c.data[user.Id] = user
	c.datas[user.Id] = statistics
	c.Unlock()
	return nil
}

func (c *Cache) Get(id uint64, token string) (*player.User, *player.Statistics, error) {
	c.Lock()
	defer c.Unlock()
	user, ok := c.data[id]

	if !ok {
		return &player.User{}, &player.Statistics{}, player.NewErrorUserNotFound(int(id))
	}
	statistic := c.datas[user.Id]

	if token != user.Token {
		return &player.User{}, &player.Statistics{}, player.NewErrorTokenNotFound(user.Id)
	}
	return user, statistic, nil
}

func(c *Cache) DepositStat(deposit *player.Deposit, user *player.User) error {
	timeDep := time.Now()
	_, ok := c.datad[deposit.DepositID]
	if ok {
		return player.NewErrorDepositIDHasBeenCreated(deposit.DepositID)
	}
	stat := player.DepositStat{
			DepositID: deposit.DepositID,
			BalanceBefore: user.Balance,
			DepositAmount: deposit.Amount,
			BalanceAfter:user.Balance + deposit.Amount,
			Time: timeDep.Format(time.RFC3339),
	}

	c.Lock()
	c.datad[deposit.DepositID] = &stat
	c.Unlock()
	return nil
}

func(c *Cache) TransactionStat(transaction *player.Transaction, user *player.User) error{
	timeDep := time.Now()
	var stat player.TransactionStat
	_, ok := c.datat[transaction.TransactionId]
	if ok {
		return player.NewErrorTransactionIDHasBeenCreated(transaction.TransactionId)
	}

	if transaction.Type == "Win" {
		stat = player.TransactionStat{
			TransactionId:     transaction.TransactionId,
			BalanceBefore:     user.Balance,
			TransactionAmount: transaction.Amount,
			BalanceAfter:      user.Balance + transaction.Amount,
			Time:              timeDep.Format(time.RFC3339),
		}
	}
	if transaction.Type == "Bet" {
		stat = player.TransactionStat{
			TransactionId:     transaction.TransactionId,
			BalanceBefore:     user.Balance,
			TransactionAmount: transaction.Amount,
			BalanceAfter:      user.Balance - transaction.Amount,
			Time:              timeDep.Format(time.RFC3339),
		}
	}
	c.Lock()
	c.datat[transaction.TransactionId] = &stat
	c.Unlock()
	return nil
}
