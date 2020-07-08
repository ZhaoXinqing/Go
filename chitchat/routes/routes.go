package routes

import (
	"net/http"
)

// 定义一个 WebRoute 结构体用于存放单个路由
type WebRoute struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// 声明 WebRoutes 切片存放所有 web 路由
type WebRoute []WebRoute

// 定义所有 Web 路由
var webRoutes = WebRoutes{
	{
		"home",
		"GET",
		"/",
		handlers.Index,
	},
	{
		"signup",
		"GET",
		"/signup"
		handlers.Signup
	},
	{
		"signupAccount",
		"POST",
		"/signip_account",
		handlers.signupAccount
	},
	{
		"login",
		"GET",
		"/login",
		handlers.Login,
	},
	{
        "auth",
        "POST",
        "/authenticate",
        handlers.Authenticate,
    },
    {
        "logout",
        "GET",
        "/logout",
        handlers.Logout,
    },
    {
        "newThread",
        "GET",
        "/thread/new",
        handlers.NewThread,
    },
    {
        "createThread",
        "POST",
        "/thread/create",
        handlers.CreateThread,
    },
    {
        "readThread",
        "GET",
        "/thread/read",
        handlers.ReadThread,
    },
    {
        "postThread",
        "POST",
        "/thread/post",
        handlers.PostThread,
    },
    {
        "error",
        "GET",
        "/err",
        handlers.Err,
    },
}