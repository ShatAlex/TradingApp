package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createTrade(c *gin.Context) {
	userId, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": userId,
	})
}

func (h *Handler) getAllTrades(c *gin.Context) {

}

func (h *Handler) getTradeById(c *gin.Context) {

}

func (h *Handler) updateTrade(c *gin.Context) {

}

func (h *Handler) deleteTrade(c *gin.Context) {

}
