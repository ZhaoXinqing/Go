// gin 框架中采用的路由库是基于httprouter做的
package webServer

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() // 默认已经连接了 Logger and Recovery 中间件
	// router := gin.New()	// 默认没有中间件的空白Gin

	// 简单的路由组：v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/sunmit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}

	// 简单的路由组：v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/sunmit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}

	// 默认在 8080 端口启动服务，除非定义了一个 PORT 的环境变量。
	router.Run()
	// router.Run(":3000") hardcode 端口号
}
