// gin 框架中采用的路由库是基于httprouter做的
package Router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello word")
	})
	r.POST("/xxxpost", getting)
	r.PUT("/xxxput")
	// 监听端口默认为8080
	r.Run(":8080")
}
