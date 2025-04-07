package model

type Goods struct {
	Id    uint64 `gorm:"column:id;type:bigint UNSIGNED;primaryKey;not null;" json:"id"`
	Name  string `gorm:"column:name;type:varchar(20);comment:商品名称;default:NULL;" json:"name"`             // 商品名称
	Price uint64 `gorm:"column:price;type:bigint UNSIGNED;comment:商品价格;not null;default:0;" json:"price"` // 商品价格
	Stock uint64 `gorm:"column:stock;type:bigint UNSIGNED;comment:商品库存;not null;default:0;" json:"stock"` // 商品库存
}

func (g *Goods) TableName() string {
	return "goods"
}

type Order struct {
	Id        uint64 `gorm:"column:id;type:bigint UNSIGNED;primaryKey;not null;" json:"id"`
	GoodsId   uint64 `gorm:"column:goods_id;type:bigint UNSIGNED;comment:商品id;not null;default:0;" json:"goods_id"`                  // 商品id
	GoodsName string `gorm:"column:goods_name;type:varchar(30);comment:用户名称;not null;" json:"goods_name"`                            // 用户名称
	Num       uint64 `gorm:"column:num;type:bigint UNSIGNED;comment:数量;not null;default:0;" json:"num"`                              // 数量
	Total     uint64 `gorm:"column:total;type:bigint UNSIGNED;comment:总金额;not null;default:0;" json:"total"`                         // 总金额
	UserName  string `gorm:"column:user_name;type:varchar(20);comment:商品名称;not null;" json:"user_name"`                              // 商品名称
	Status    uint8  `gorm:"column:status;type:tinyint UNSIGNED;comment:订单状态 0:待支付 1:支付成功 2：取消支付;not null;default:0;" json:"status"` // 订单状态 0:待支付 1:支付成功 2：取消支付
}

func (o *Order) TableName() string {
	return "order"
}

type User struct {
	Id       uint64 `gorm:"column:id;type:bigint UNSIGNED;primaryKey;not null;" json:"id"`
	Username string `gorm:"column:username;type:varchar(30);comment:用户名;not null;" json:"username"`            // 用户名
	Password string `gorm:"column:password;type:char(32);comment:密码;not null;" json:"password"`                // 密码
	Balance  uint64 `gorm:"column:balance;type:bigint UNSIGNED;comment:余额;not null;default:0;" json:"balance"` // 余额
}

func (u *User) TableName() string {
	return "user"
}
