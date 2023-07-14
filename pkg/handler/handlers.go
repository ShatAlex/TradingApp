package handler

import (
	"github.com/ShatAlex/trading-app/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(ser *service.Service) *Handler {
	return &Handler{services: ser}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api/v1", h.userIdentity)
	{
		treasuries := api.Group("/treasuries")
		{
			treasuries.POST("/", h.getTickerPrice)
		}
		trades := api.Group("/trades")
		{
			trades.POST("/", h.createTrade)
			trades.GET("/", h.getAllTrades)
			trades.GET("/:id", h.getTradeById)
			trades.PUT("/:id", h.updateTrade)
			trades.DELETE("/:id", h.deleteTrade)

			trades.POST("/buy", h.buyTicker)
			trades.POST("/sell", h.sellTicker)
			trades.GET("/portfolio", h.getPortfolio)
		}
		types := api.Group("/types")
		{
			types.POST("/", h.createType)
			types.GET("/", h.getAllType)
			types.GET("/:id", h.getTypeById)
			types.PUT("/:id", h.updateType)
			types.DELETE("/:id", h.deleteType)
		}
	}
	return router
}
