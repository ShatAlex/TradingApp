package handler

import (
	"net/http"

	trade "github.com/ShatAlex/trading-app"
	"github.com/gin-gonic/gin"
)

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

	id, err := h.services.Portfolio.BuyTicker(userId, input, ticker[0].Close)

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

	total, err := h.services.Portfolio.SellTicker(userId, input, ticker[0].Close)

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

	tickers, err := h.services.Portfolio.GetAllTickers(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tickers)
}

func (h *Handler) getSpecificTicker(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	nasdaq := c.Query("ticker")

	ticker, err := h.services.Portfolio.GetTickerByNasdaq(userId, nasdaq)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, ticker)
}
