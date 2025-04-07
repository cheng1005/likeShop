package main

import (
	"github.com/gin-gonic/gin"
	_ "zg4/likeShop/api-gateway/basic/init"
	"zg4/likeShop/api-gateway/router"
)

func main() {
	r := gin.Default()

	router.LoadRouter(r)

	r.Run()
}
