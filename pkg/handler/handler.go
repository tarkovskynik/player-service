package handler

import (
	"github.com/gin-gonic/gin"
	"player/pkg/cache"
	"player/pkg/database"
)

type Handler struct {
	repo  *database.PlayerRepository
	cache *cache.Cache
}

func NewHandler(repo *database.PlayerRepository, cache *cache.Cache) *Handler {
	return &Handler{
		repo:  repo,
		cache: cache,
	}

}

func (h *Handler) Init() error {
	r := gin.New()

	r.POST("/user/create", h.createUser)
	r.GET("/user/get", h.getUser)
	r.POST("/user/deposit", h.addDeposit)
	r.POST("/transaction", h.transaction)

	return r.Run(":8080")
}
