package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
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
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid authorization header")
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	c.Set(userCtx, userId)
}
