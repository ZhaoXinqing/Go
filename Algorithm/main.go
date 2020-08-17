func main() {
	r := gin.Default()

	r.GET("/someJSON",func(c *gin.Context){
		data := map[string]interface{} {
			"lang":"Go语言",
			"tag":"<br>",
		}
		c.AsciiJSON(http.StatusOK,data)
	})

	r.Run(":8080")
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/index",func(c *gin.Context) {
		c.HTML(htp.StatusOK,"index.tmpl",gin.H{
			"title":"Main website",
		})
	})
	router.Run(":8080")
}

var html = template.Must(template.New("https").Parse(

))

func main() {
	r := gin.Default()
	r.Static("/assets","./assets")
	r.SetHTMLTemplate(html)

	r.GET("/",)
}