package handler

import (
	"net/http"
	"strconv"

	trade "github.com/ShatAlex/trading-app"
	"github.com/gin-gonic/gin"
)

// @Summary Create trade
// @Security ApiKeyAuth
// @Tags Trades
// @Description create trade
// @ID create-trade
// @Accept  json
// @Produce  json
// @Param input body trade.Trade true "trade fields"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/trades [post]
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

// @Summary Get trades
// @Security ApiKeyAuth
// @Tags Trades
// @Description get all trades
// @ID get-trades
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/trades [get]
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

// @Summary Get specific trade
// @Security ApiKeyAuth
// @Tags Trades
// @Description get specific trade
// @ID get-specific-trade
// @Accept  json
// @Produce  json
// @Param id path int true "Trade ID"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/trades/{id} [get]
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

// @Summary Update trade
// @Security ApiKeyAuth
// @Tags Trades
// @Description update trade
// @ID update-trade
// @Accept  json
// @Produce  json
// @Param input body trade.Trade true "trade fields"
// @Param id path int true "Trade ID"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/trades/{id} [put]
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

// @Summary Delete trade
// @Security ApiKeyAuth
// @Tags Trades
// @Description delete trade
// @ID delete-trade
// @Accept  json
// @Produce  json
// @Param id path int true "Trade ID"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/trades/{id} [delete]
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
