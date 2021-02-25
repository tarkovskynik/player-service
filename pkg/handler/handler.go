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

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"admin": "qwerty",
	}))

	authorized.POST("/user/create", h.createUser)
	authorized.GET("/user/get", h.getUser)
	authorized.POST("/user/deposit", h.addDeposit)
	authorized.POST("/transaction", h.transaction)

	return r.Run(":8080")
}
