package cache

import (
	"player"
	"sync"
)

type Cache struct {
	data  map[uint64]*player.User
	datas map[uint64]*player.Statistics
	datad map[uint64]*player.Deposit
	datat map[uint64]*player.Transaction
	sync.Mutex
}

func NewCache() *Cache {
	return &Cache{
		data:  make(map[uint64]*player.User),
		datas: make(map[uint64]*player.Statistics),
		datad: make(map[uint64]*player.Deposit),
		datat: make(map[uint64]*player.Transaction),
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

func (c *Cache) UpdateCache() {

}
