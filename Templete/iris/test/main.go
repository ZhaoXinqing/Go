package main

import (
	"github.com/kataras/iris"
    "github.com/kataras/iris/middleware/basicauth"
)

func newApp() *iris.Application {
	app := iris.New()

	authConfig := basicauth.Config{
		User: map[string]string{"myusername":"mypassword"},
	}

	authentication := basicauth.New(authConfig)
	app.Get("/",func(ctx iris.Context){ ctx.Redirect("/admin")})
	needAuth := app.Party("/admin",authentication)
	{
		needAuth.Get("/",h)
		needAuth.Get("/profile",h)
		needAuth.Get("/settings",h)
	}

	return app
}

type User struct { Username string }
func (u *User) Decode(data []byte) error {
	return json.Unmarshal(data,u)
}