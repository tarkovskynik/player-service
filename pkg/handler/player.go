package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"player"
)

type errorResponse struct {
	Error string `json:"error"`
}

func (h *Handler) createUser(c *gin.Context) {
	var input player.User
	var stat player.Statistics

	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.WithField("handler", "createUser").Errorf("error: %s", err.Error())
		c.JSON(http.StatusBadRequest, errorResponse{
			Error: err.Error(),
		})
		return
	}

//	id, err := h.repo.Create(input)
    err := h.cache.Create(&input, &stat)
    if err != nil {
		logrus.WithField("handler", "createUser").Errorf("error: %s", err.Error())
		c.JSON(http.StatusBadRequest, errorResponse{
			Error: err.Error(),
		})
		return
	}


	c.JSON(http.StatusOK, map[string]interface{}{
		"id": input.Id,
	})
}

func (h *Handler) getUser(c *gin.Context) {
	var input player.User
	//var statistics player.Statistics

	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.WithField("handler", "getUser").Errorf("error: %s", err.Error())
		c.JSON(http.StatusBadRequest, errorResponse{
			Error: err.Error(),
		})
		return
	}
	user, statistic, err := h.cache.Get(input.Id,input.Token)

		if err != nil {
			logrus.WithField("handler", "getUser").Errorf("error: %s", err.Error())
			c.JSON(http.StatusInternalServerError, errorResponse{
				Error: err.Error(),
			})
			return
		}

	c.JSON(http.StatusOK, user)
	c.JSON(http.StatusOK, statistic)
}

func (h *Handler) addDeposit(c *gin.Context) {
	var deposit player.Deposit

	if err := c.ShouldBindJSON(&deposit); err != nil {
		logrus.WithField("handler", "getUser").Errorf("error: %s", err.Error())
		c.JSON(http.StatusBadRequest, errorResponse{
			Error: err.Error(),
		})
		return
	}
	user, statistics, err := h.cache.Get(deposit.UserID,deposit.Token)

	if err != nil {
		logrus.WithField("handler", "getUser").Errorf("error: %s", err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse{
			Error: err.Error(),
		})
		return
	}
	statistics.DepositCount++
	statistics.DepositSum += deposit.Amount
	user.Balance += deposit.Amount

	c.JSON(http.StatusOK, user.Balance)
}

func (h *Handler) transaction(c *gin.Context){
	var transaction player.Transaction

	if err := c.ShouldBindJSON(&transaction); err != nil {
		logrus.WithField("handler", "getUser").Errorf("error: %s", err.Error())
		c.JSON(http.StatusBadRequest, errorResponse{
			Error: err.Error(),
		})
		return
	}

	user, statistics, err := h.cache.Get(transaction.UserID,transaction.Token)

	if err != nil {
		logrus.WithField("handler", "getUser").Errorf("error: %s", err.Error())
		c.JSON(http.StatusInternalServerError, errorResponse{
			Error: err.Error(),
		})
		return
	}
	if transaction.Type == "Win" {
		statistics.WinCount++
		statistics.WinSum += transaction.Amount
		user.Balance += transaction.Amount
	}
	if transaction.Type == "Bet" {
		statistics.BetCount++
		statistics.BetSum += transaction.Amount
		user.Balance -= transaction.Amount
	}


	c.JSON(http.StatusOK, user.Balance)
}