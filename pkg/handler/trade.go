package handler

import (
	"net/http"
	"strconv"

	trade "github.com/ShatAlex/trading-app"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createTrade(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input trade.Trade
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Trade.Create(userId, input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllTradesResponse struct {
	Data []trade.Trade `json:"data"`
}

func (h *Handler) getAllTrades(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	trades, err := h.services.Trade.GetAll(userId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllTradesResponse{
		Data: trades,
	})
}

func (h *Handler) getTradeById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	tradeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid tradeId")
		return
	}

	item, err := h.services.Trade.GetTradeById(userId, tradeId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) updateTrade(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	tradeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid tradeId")
		return
	}

	var input trade.UpdateTradeInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Trade.Update(userId, tradeId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) deleteTrade(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	tradeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid tradeId")
		return
	}

	if err := h.services.Trade.Delete(userId, tradeId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) buyTicker(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input trade.BuySellTickerInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.Amount == nil {
		newErrorResponse(c, http.StatusBadRequest, "empty amount field")
		return
	}

	ticker, err := getTreasuries(c, input.Ticker)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Trade.BuyTicker(userId, input, ticker[0].Close)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) sellTicker(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input trade.BuySellTickerInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.Amount == nil {
		newErrorResponse(c, http.StatusBadRequest, "empty amount field")
		return
	}

	ticker, err := getTreasuries(c, input.Ticker)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	total, err := h.services.Trade.SellTicker(userId, input, ticker[0].Close)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"total": total,
	})
}

func (h *Handler) getPortfolio(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	tickers, err := h.services.Trade.GetAllTickers(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tickers)
}
