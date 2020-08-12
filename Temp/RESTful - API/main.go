package main

import (
	"io"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-swagger12"
)

func main() {
	wsContainer := restful.NewContainer()

	// 跨域过滤器
	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{"X-My-Header"},
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"GET", "POST"},
		CookiesAllowed: false,
		Container:      wsContainer}
	wsContainer.Filter(cors.Filter)

	// Add container filter to respond to OPTIONS
	wsContainer.Filter(wsContainer.OPTIONSFilter)

	config := swagger.Config{
		WebServices:    restful.DefaultContainer.RegisteredWebServices(), // you control what services are visible
		WebServicesUrl: "http://localhost:8080",
		ApiPath:        "/apidocs.json",
		ApiVersion:     "V1.0",
		// Optionally, specify where the UI is located
		SwaggerPath:     "/apidocs/",
		SwaggerFilePath: "D:/gowork/src/doublegao/experiment/restful/dist"}
	swagger.RegisterSwaggerService(config, wsContainer)
	//swagger.InstallSwaggerService(config)

	u := UserResource{}
	u.RegisterTo(wsContainer)

	log.Print("start listening on localhost:8080")
	server := &http.Server{Addr: ":8080", Handler: wsContainer}
	defer server.Close()
	log.Fatal(server.ListenAndServe())

}

type UserResource struct{}

func (u UserResource) RegisterTo(container *restful.Container) {
	ws := new(restful.WebService)
	//设置匹配的schema和路径
	ws.Path("/user").Consumes("*/*").Produces("*/*")

	//设置不同method对应的方法，参数以及参数描述和类型
	//参数:分为路径上的参数,query层面的参数,Header中的参数
	ws.Route(ws.GET("/{id}").
		To(u.result).
		Doc("方法描述：获取用户").
		Param(ws.PathParameter("id", "参数描述:用户ID").DataType("string")).
		Param(ws.QueryParameter("name", "用户名称").DataType("string")).
		Param(ws.HeaderParameter("token", "访问令牌").DataType("string")).
		Do(returns200, returns500))
	ws.Route(ws.POST("").To(u.result))
	ws.Route(ws.PUT("/{id}").To(u.result))
	ws.Route(ws.DELETE("/{id}").To(u.result))

	container.Add(ws)
}

func (UserResource) SwaggerDoc() map[string]string {
	return map[string]string{
		"":         "Address doc", //空表示结构本省的描述
		"country":  "Country doc",
		"postcode": "PostCode doc",
	}
}

func (u UserResource) result(request *restful.Request, response *restful.Response) {
	io.WriteString(response.ResponseWriter, "this would be a normal response")
}

func returns200(b *restful.RouteBuilder) {
	b.Returns(http.StatusOK, "OK", "success")
}

func returns500(b *restful.RouteBuilder) {
	b.Returns(http.StatusInternalServerError, "Bummer, something went wrong", nil)
}
