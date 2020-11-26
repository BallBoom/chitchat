package router

import (
	"main/handlers"
	"net/http"
)

// 定义路由结构体  结构体用于存放单个路由
type WebRoute struct {
	Name        string
	Methods     string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// 声明WebRoutes，切片存放所有web 路由
type WebRoutes []WebRoute

// 定义所有 Web 路由
var webRoutes = WebRoutes{
	{
		"home",
		"GET",
		"/",
		handlers.Index,
	},
	{
		"login",
		"GET",
		"/login",
		handlers.Login,
	},
	{
		"logout",
		"GET",
		"/logout",
		handlers.Logout,
	},
	{
		"auth",
		"POST",
		"/authenticate",
		handlers.Authenticate,
	},
	{
		"singup",
		"GET",
		"/signup",
		handlers.Signup,
	},
	{
		"signup_account",
		"POST",
		"/signup_account",
		handlers.SingupAccount,
	},
}
