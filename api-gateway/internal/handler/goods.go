package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zg4/likeShop/api-gateway/global"
	"zg4/likeShop/api-gateway/internal/request"
	"zg4/likeShop/api-gateway/proto/goods"
)

func AddOrder(c *gin.Context) {

	var req request.AddOrderReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	order, err := global.GoodsClient.AddOrder(c, &goods.AddOrderRequest{
		GoodsId: req.GoodsId,
		Num:     req.Num,
		UserId:  req.UserId,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "创建订单成功",
		"data": order,
	})

}

func OrderList(c *gin.Context) {
	var req request.OrderListReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	orderList, err := global.GoodsClient.OrderList(c, &goods.OrderListRequest{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "订单列表展示成功",
		"data": orderList,
	})
}
