package handler

import (
	"context"
	"log"
	"net/http"

	trade "github.com/ShatAlex/trading-app"
	"github.com/gin-gonic/gin"
	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

func (h *Handler) getTreasuries(c *gin.Context) {

	var input trade.PolygonInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	pol := polygon.New("4Sag3Fh4lQIYMt6P219ziTdI_nT1Pnec")

	params := models.GetPreviousCloseAggParams{
		Ticker: *input.Ticker,
	}

	res, err := pol.GetPreviousCloseAgg(context.Background(), &params)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ticker": res.Results[0].Ticker,
		"price":  res.Results[0].Close,
	})
}
