package main

import (
	"fmt"
	"github.com/kataras/iris"
)

func main() {
	fmt.Printf("Hello, zhaoxinqing")
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {})
	app.Run(iris.Addr(":8080"))
}
