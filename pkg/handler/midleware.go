package handler

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

const (
	autorizationHeader = "Authorization"
	userCtx            = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(autorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty authorization header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, "invalid authorization header")
		return
	}

	if len(headerParts[1]) == 0 {
		newErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "userId not found")
		return 0, errors.New("userId not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "ivalid type userId")
		return 0, errors.New("invalid type userId")
	}

	return idInt, nil
}

func getTreasuries(c *gin.Context, ticker *string) ([]models.Agg, error) {

	pol := polygon.New("4Sag3Fh4lQIYMt6P219ziTdI_nT1Pnec")

	params := models.GetPreviousCloseAggParams{
		Ticker: *ticker,
	}

	res, err := pol.GetPreviousCloseAgg(context.Background(), &params)
	if err != nil {
		return res.Results, errors.New("bad request")
	}

	return res.Results, nil
}
