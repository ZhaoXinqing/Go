package main

import(
	"github.com/gin-gonic/gin"
	"file-strong/handler"
)

func main() {
	router := gin.defer Default()
	router.GET("/ping",func(c * gin.Context){
		c.JSON(200,gin.H{
			"message":"pong",
		})
	})

	// router.GET("/someGet", GetHandler)
	router.POST("/post",handler.PostHandler)

	// 默认启动的是8080端口，也可以自己定义启动端口
	router.Run()
}