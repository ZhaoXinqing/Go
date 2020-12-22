package webServer

//import (
//	"net/http"
//
//	"github.com/gin-gonic/gin"
//)
//
//func main() {
//	// 1.创建路由
//	r := gin.Default() // 默认已经连接了 Logger and Recovery 中间件
//	// r := gin.New()	  // 默认没有中间件的空白Gin
//
//	// 2.绑定路由的规则，执行的函数
//	// gin.Context.封装了request和response
//	r.GET("/", func(c *gin.Context) {
//		c.String(http.StatusOK, "hello World!")
//	})
//	// 3.监听端口，默认在8080
//	// Run("里面不指定端口号默认为8080")
//	r.Run(":8080")
//}
