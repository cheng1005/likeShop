package init

import (
	"fmt"
	"github.com/spf13/viper"

	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net"
	"time"
	"zg4/likeShop/goods-server/basic/config"
	"zg4/likeShop/goods-server/global"
	"zg4/likeShop/goods-server/internal/handler"
	"zg4/likeShop/goods-server/proto/goods"
)

func init() {

	InitConfig()
	InitMysql()
	InitGrpc()
}

func InitMysql() {
	var err error
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Con.Mysql.User, config.Con.Mysql.Password, config.Con.Mysql.Host, config.Con.Mysql.Port, config.Con.Mysql.Database)
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := global.DB.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("mysql init success")
}

func InitGrpc() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8070))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	goods.RegisterGoodsSrvServer(s, &handler.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func InitConfig() {

	viper.SetConfigFile("./dev.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&config.Con)
	if err != nil {
		panic(err)
	}
	log.Println("config read success")
}
