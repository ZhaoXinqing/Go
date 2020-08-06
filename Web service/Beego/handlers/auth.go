package handlers

import (
    "github.com/xueyuanjun/chitchat/models"
    "net/http"
)

// Get/login
// 登陆界面
func login(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "auth.layout", "navbar", "login")
}

// GET/signup
// 注册页面
func Signup(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "auth.layout", "navbar", "signup")
}

// POST /signup
// 注册页面
func SignupAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		danger(err, "Cannot prase form")
	}
	user := models.User{
		Name:request.PostFormValue("name")
		Email: request.PostFormValue("email")
		Password: request.PostFormValue("password")
	}
	if err := user.Create();err != nil() {
		danger(err,"Cannot creat user")
	}
	http.Redirect(writer, request,"/login",302)
}

// POST /authenticate
// 通过邮箱和密码字段对用户进行认证
func Authenticate(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	user, err := models.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		danger(err, "Cannot find user")
	}
	if user.Password == models.Encrypt(request.PostFormValue("password")) {
		session,err := user.CreateSession()
		if err != nil {
			danger(err, "Cannot creat session")
		}
		cookie := http.Cookie{
			Name: "_cookie",
			Value: session.Uuid,
			HttpOnly:true,
		}
		http.SetCookie(writer,&cookie)
		http.Redirect(writer,request,"/", 302)
	} else {
		http.Redirect(writer,request,"/login",302)
	}
}

// GET /login
// 用户退出
func logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != http.ErrNoCookie {
		warning(err, "Failed to get cookie")
		session := models.session{Uuid: cookie.Value}
		session.DeleteByUUID()
	}
	http.Redirect(writer,request,"/", 302)
}
