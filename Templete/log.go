package main

import (
	"os"
	"time"

	"github.com/kataras/iris"
)

func todayFilename() string {
	today := time.Now().Format("Jan 02 2006")
	return today + ".txt"
}

func newLogFile() *os.File {
	filename := todayFilename()

	f, err := os.OpenFile(filename,os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err != nil {
		panic(err)
	}
	return f
}

func main() {
	f := newLogFile()
	defer f.Close()

	app := iris.New()
	app.Logger().SetOutput(f)

	app.Get("/ping",func(ctx iris.Context) {
		ctx.Application().Logger().Info("Request path:%s",ctx.Path())
		ctx.WriteString("pong")
	})

	app.Run(
		iris.Addr(":8080"),
		iris.WithoutBanner,
		iris.WithoutServerError(iris.ErrServerClosed),
	)
}

// 删除 Cookie.
app.Delete("/cookie/{name}",func(ctx iris.Context) {
	name := ctx.Params().Get("name")
	ctx.RemoveCookie(name)
	ctx.Writef("cookie %s remove", name)
})

return app

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}