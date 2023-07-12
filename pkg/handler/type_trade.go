package handler

import (
	"net/http"
	"strconv"

	trade "github.com/ShatAlex/trading-app"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createType(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input trade.TypeTrade
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TypeTrade.Create(userId, input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllListsResponse struct {
	Data []trade.TypeTrade `json:"data"`
}

func (h *Handler) getAllType(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	types, err := h.services.TypeTrade.GetAll(userId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: types,
	})
}

func (h *Handler) getTypeById(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	typeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid typeId")
		return
	}

	item, err := h.services.TypeTrade.GetTypeById(userId, typeId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) updateType(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	typeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid typeId")
		return
	}

	var input trade.TypeTrade
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TypeTrade.Update(userId, typeId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})

}

func (h *Handler) deleteType(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	typeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid typeId")
		return
	}

	if err := h.services.TypeTrade.Delete(userId, typeId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
