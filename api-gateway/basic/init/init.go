package init

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"zg4/likeShop/api-gateway/global"
	"zg4/likeShop/api-gateway/internal/untils"
	"zg4/likeShop/api-gateway/proto/goods"
)

func init() {
	
	InitConsul()
	InitGoodsClient()
}

func InitZap() {

	untils.InitLogger()
	defer untils.SugarLogger.Sync()
	untils.SimpleHttpGet("./likeShop.log")

}

func InitConsul() {

	client, err := untils.NewConsulClient("14.103.235.216", 8500)
	if err != nil {
		return
	}
	register := untils.ConsulRegister{
		Id:      "1314",
		Name:    "xxfz",
		Tags:    []string{"zk2"},
		Port:    8001,
		Address: "localhost",
	}
	err = client.Register(register)
	if err != nil {
		return
	}
	log.Println("consul register success")
}

func InitGoodsClient() {

	goodsConn, err := grpc.NewClient("localhost:8070", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	global.GoodsClient = goods.NewGoodsSrvClient(goodsConn)
}
