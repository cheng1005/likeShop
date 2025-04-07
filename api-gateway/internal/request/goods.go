package request

// AddOrderReq 创建订单请求参数
type AddOrderReq struct {
	GoodsId uint64 `form:"goods_id" json:"goods_id" binding:"required"`
	UserId  uint64 `form:"user_id" json:"user_id" binding:"required"`
	Num     uint64 `form:"num" json:"num" binding:"required"`
}

type OrderListReq struct {
	Page uint64 `form:"page" json:"page"`
	Size uint64 `form:"size" json:"size"`
}
