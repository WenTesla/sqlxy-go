package main

import (
	"Project/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	// 路由
	group := engine.Group("/discuss")
	group.PUT("/action", controller.DiscussActionController)
	group.POST("/list", controller.DiscussListController)
	engine.Run(":8080")

}
