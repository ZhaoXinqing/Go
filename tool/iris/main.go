package main

import (
	""
)

func main() {
	app := iris.Default()

	app.Get("/someGet",posting)
	app.Post("/somePost",posting)
	app.Put("/somePut",putting)
	app.Delete("/someDelete",delete)
	app.Patch("/somePatch",patching)
	app.Head("/someHead",head)
	app.Options("/someOptions",Options)

	app.Run(iris.Addr(":8080"))
}

// ctx 实现上下文接口类型实例，其内持有 net/http 请求与响应实例，实质是一个组合对象。
app.Get("/users/{id:uint64}",func(ctx iris.Context){
	id := ctx.Params().GetUint64Default("id",0)
	// [...]
})

	app.Post("/user/{name:string}/{action:path}",func(ctx iris.Context){
		name := ctx.Params().Get("name")
		action := Params().Get("action")
		message := name + " is " + action
		ctx.WriteString(message)
	})

	app.Run(iris(":8080"))
}

// 发起一个 post 请求
func main() {
	app := iris.Default()

	app.Post("/form_post",func(ctx iris.Context) {
		message := ctx.FormValue("message")
		nick  := ctx.FormValueDefault("nick","anonymous")
		ctx.JSON(iris.Map{
			"status": "posted",
			"message": message,
			"nick": nick,
		})
	})
}

// 表单
func main() {
	app := iris.Default()

	app.Post("/post",func(ctx iris.context){
		id := ctx.URLParam("id")
		page := ctx.URLParamDefault("page","0")
		name := ctx.FormValue("name")
		message := ctx.FormValue("message")

		app.Logger().Infof("id:%s; page: %s; name: %s; message: %s", id, page, name,message)
	})

	app.Run(iris.Addr(":8080"))
}

// 文件上传
sonst maxSize = 5 << 20

func main() {
	app := iris.Default()
	app.Post("/upload",iris.LimitRequestBodySize(maxSize),func(ctxiris.Context){
		ctx.uploadFormFiles("./uploads", beforeSave)
	})

	app.Run(iris.Addr(":8080"))
}

func beforeSave(ctx iris.Context,file *multipart.FileHeader) {
	ip := ctx.RemoteAddr()
	ip = strings.Replace(ip, ".","_",-1)
	ip = strings.Replace(ip, ":","_",-1)

	file.Filename = ip + "-" + file.Filename
}

// 命令行请求
curl -X POST http://localhost:8080/upload \
	-F "files[]=@./myfiile.zip" \
	-F "files[]=@./mysecondfile.zip" \
	-H "Content-Type:multipart/form-data"


func main() {
	app := iris.New()

	app.Get("/",func(ctx content.Context))
}

// 路由分组
func main() {
	app := iris.Default()

	// 简单分组：v1
	v1 := app.Party("/v1")
	{
		v1.Post("/login",loginEndpoint)
		v1.Post("/submit",submitEndpoint)
		v1.Post("/read",readEndpoint)
	}

	v2 := app.Party("/v2")
	{
		v2.Post("/login",loginEndpoint)
		v2.Post("/submit",submitEndpoint)
		v2.Post("/read",readEndpoint)
	}

	app.Run(iris.Addr(":8080"))
}


app := iris.New()

// Recovery
app := iris.Default()

func main() {
	app := iris.New()
	app.Use(recover.New())

	requestLogger := logger.New(logger.Config{
		Status: true,
		IP:true,
		Method: true,
		Path:true,
		Query:true,

		authorized := app.Party("/user")
		authorized.Use(AuthRequired())
		{
			authorized.Post("/login",loginEndpoint)
			authorized.Post("/submit",submitEndpoint)
			authorized.Post("/read",readEndpoint)

			testing := authorized.Party("testing")
			testing.Get("/analytics",analyticsEndpoint)
		}
	})
}

