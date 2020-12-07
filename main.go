package main

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "usercenter/docs"
	"usercenter/router"
)

// @title 用户中心API
// @version 1.0
func main() {
	engine := router.SetupRouter()

	url := ginSwagger.URL("./doc.json")
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	_ = engine.Run(":8081")
}
