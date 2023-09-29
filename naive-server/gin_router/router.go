package gin_router

import (
	api "qsc/gin_api"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(gin.ReleaseMode)
	//使用gin的Default方法创建一个路由handler
	router := gin.Default()
	router.GET("/ping", api.Ping)
	router.POST("/signup", api.Signup)
	router.POST("/signin", api.Signin)
	router.POST("/checkin", api.Checkin)
	//监听服务器端口
	router.Run(":8080")
}
