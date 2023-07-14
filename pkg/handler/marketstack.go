package handler

import (
	"net/http"

	trade "github.com/ShatAlex/trading-app"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getTickerPrice(c *gin.Context) {

	var input trade.PolygonInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := getTreasuries(c, input.Ticker)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ticker": res[0].Ticker,
		"price":  res[0].Close,
	})
}
