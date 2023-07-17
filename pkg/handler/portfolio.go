package handler

import (
	"net/http"

	trade "github.com/ShatAlex/trading-app"
	"github.com/gin-gonic/gin"
)

// @Summary Buy ticker
// @Security ApiKeyAuth
// @Tags Portfolio
// @Description buy ticker
// @ID buy-ticker
// @Accept  json
// @Produce  json
// @Param input body trade.BuySellTickerInput true "ticker fields"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/portfolio/buy [post]
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

// @Summary Sell ticker
// @Security ApiKeyAuth
// @Tags Portfolio
// @Description sell ticker
// @ID sell-ticker
// @Accept  json
// @Produce  json
// @Param input body trade.BuySellTickerInput true "ticker fields"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/portfolio/sell [post]
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

// @Summary Get portfolio
// @Security ApiKeyAuth
// @Tags Portfolio
// @Description get portfolio
// @ID get-portfolio
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/portfolio [get]
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

// @Summary Get specific ticker
// @Security ApiKeyAuth
// @Tags Portfolio
// @Description get specific ticker
// @ID get-cpecific-tciker
// @Accept  json
// @Produce  json
// @Param ticker path string true "Ticker NASDAQ"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/portfolio/detail?ticker={ticker} [get]
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

// @Summary Get income
// @Security ApiKeyAuth
// @Tags Portfolio
// @Description get income
// @ID get-income
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/portfolio/income [get]
func (h *Handler) getIncome(c *gin.Context) {

	var total float64

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	portfolio, err := h.services.Portfolio.GetAllTickers(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	for _, elem := range portfolio {
		ticker, err := getTreasuries(c, &elem.Ticker)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		total += ticker[0].Close * float64(elem.Amount)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"total_income": total,
	})

}
