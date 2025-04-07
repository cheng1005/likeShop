package handler

import (
	"context"
	"fmt"
	"zg4/likeShop/goods-server/global"
	"zg4/likeShop/goods-server/internal/model"
	"zg4/likeShop/goods-server/pkg"
	"zg4/likeShop/goods-server/proto/goods"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	goods.UnimplementedGoodsSrvServer
}

func (s *Server) AddOrder(_ context.Context, req *goods.AddOrderRequest) (*goods.AddOrderResponse, error) {

	var good model.Goods
	global.DB.Debug().Where("id = ?", req.GoodsId).Find(&good)
	if good.Id == 0 {
		return nil, fmt.Errorf("商品不存在")
	}
	if good.Stock < req.Num {
		return nil, fmt.Errorf("商品库存不足")
	}
	var user model.User
	global.DB.Debug().Where("id = ?", req.UserId).Find(&user)
	if user.Id == 0 {
		return nil, fmt.Errorf("用户信息异常")
	}
	total := good.Price * req.Num
	order := model.Order{
		GoodsId:   good.Id,
		GoodsName: good.Name,
		Num:       req.Num,
		Total:     total,
		UserName:  user.Username,
		Status:    0,
	}
	if err := global.DB.Debug().Create(&order).Error; err != nil {
		return nil, fmt.Errorf("订单创建失败")
	}
	payUrl := pkg.AliPay(good.Name, int(total))
	return &goods.AddOrderResponse{
		PayUrl: payUrl,
	}, nil
}

func (s *Server) OrderList(_ context.Context, req *goods.OrderListRequest) (*goods.OrderListResponse, error) {
	page := req.Page
	if page <= 0 {
		page = 1
	}

	pageSize := req.Size
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	var orderList []model.Order
	global.DB.Debug().Offset(int(offset)).Limit(int(pageSize)).Find(&orderList)

	var lists []*goods.OrderList
	for _, v := range orderList {
		lists = append(lists, &goods.OrderList{
			OrderId:   v.Id,
			GoodsId:   v.GoodsId,
			GoodsName: v.GoodsName,
			Num:       v.Num,
			Total:     v.Total,
			UserName:  v.UserName,
			Status:    uint64(v.Status),
		})
	}
	return &goods.OrderListResponse{
		List: lists,
	}, nil
}
