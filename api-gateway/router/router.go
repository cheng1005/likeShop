package router

import (
	"github.com/gin-gonic/gin"
	"zg4/likeShop/api-gateway/internal/handler"
)

func LoadRouter(r *gin.Engine) {
	apiRouter := r.Group("/api")
	{
		goods := apiRouter.Group("/goods")
		{
			goods.POST("/add-order", handler.AddOrder)
			goods.GET("/order-list", handler.OrderList)

		}
	}
}
