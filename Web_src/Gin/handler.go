import (
	"github.com/gin-gonic/gin"
)

func PostHandler(c *gin.Context) {
	chk_id := c.PostForm("chk_id")
	chk_id2 := c.Query("chk_id")
	panth := c.DefaultPostForm("path", "anonymous") // 此方法可以设置默认值

	c.JSON(200, gin.H{
		"status":  "posted",
		"message": "ok",
		"path":    "path",
		"chk_id":  chk_id,
		"chk_id2": chk_id2,
	})
}