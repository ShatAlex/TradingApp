package handler

import (
	"net/http"
	"strconv"

	trade "github.com/ShatAlex/trading-app"
	"github.com/gin-gonic/gin"
)

// @Summary Create type
// @Security ApiKeyAuth
// @Tags Types
// @Description create type
// @ID create-type
// @Accept  json
// @Produce  json
// @Param input body trade.TypeTrade true "type fields"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/types [post]
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

type getAllTypesResponse struct {
	Data []trade.TypeTrade `json:"data"`
}

// @Summary Get types
// @Security ApiKeyAuth
// @Tags Types
// @Description getting the type
// @ID get-type
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/types [get]
func (h *Handler) getAllType(c *gin.Context) {

	types, err := h.services.TypeTrade.GetAll()

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllTypesResponse{
		Data: types,
	})
}

// @Summary Get specific type
// @Security ApiKeyAuth
// @Tags Types
// @Description getting specific type
// @ID get-specific-type
// @Accept  json
// @Produce  json
// @Param id path int true "Type ID"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/types/{id} [get]
func (h *Handler) getTypeById(c *gin.Context) {

	typeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid typeId")
		return
	}

	item, err := h.services.TypeTrade.GetTypeById(typeId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

// @Summary Update type
// @Security ApiKeyAuth
// @Tags Types
// @Description update the type
// @ID update-type
// @Accept  json
// @Produce  json
// @Param id path int true "Type ID"
// @Param input body trade.UpdateTradeInput true "type fildes"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/types/{id} [put]
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

// @Summary Delete type
// @Security ApiKeyAuth
// @Tags Types
// @Description delete type
// @ID delete-type
// @Accept  json
// @Produce  json
// @Param id path int true "Type ID"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/types/{id} [delete]
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
