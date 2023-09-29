// package main

// import (
// 	"net/http"
// 	"qsc/common"

// 	"github.com/gin-gonic/gin"
// )

//	func main() {
//		r := gin.Default()
//		common.InitDB()
//		db := common.GetDB()
//		defer db.Close()
//		r.GET("/ping", func(c *gin.Context) {
//			c.JSON(http.StatusOK, gin.H{
//				"code": 200,
//				"msg":  "pong",
//			})
//		})
//		r.Run(":8080")
//	}
package main

import (
	Main "qsc/gin_router"
)

func main() {
	Main.InitRouter()
}
