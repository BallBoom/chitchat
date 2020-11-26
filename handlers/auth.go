package handlers

import (
	"fmt"
	"main/moudels"
	"net/http"
)

// GET /login
// 登陆
func Login(w http.ResponseWriter, r *http.Request) {
	t := parseTemplateFiles("auth.layout", "navbar", "login")
	t.Execute(w, nil)
}

// GET /singup
// 注册页面
func Signup(w http.ResponseWriter, r *http.Request) {
	generateHtml(w, nil, "auth.layout", "navbar", "signup")
}

// POST /singup
//注册新用户
func SingupAccount(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		// log.Fatal(err.Error())
		fmt.Println(err.Error())

	}

	user := moudels.User{
		Name:     r.PostFormValue("name"),
		Password: r.PostFormValue("password"),
		Email:    r.PostFormValue("email"),
	}
	if _, err := user.Create(); err != nil {
		//log.Fatal("connot create user")
		fmt.Println("connot create user")
	}
	http.Redirect(w, r, "/login", 302)
}

//POST /authenticate
// 通过邮箱和密码验证用户 登陆
func Authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		//log.Fatal(err.Error())
		fmt.Println(err.Error())
	}
	email := r.PostFormValue("email")
	user, err := moudels.UserByEmail(email)
	if err != nil {
		//log.Fatal("connot find user")
		fmt.Println("connot find user")
	}
	password := r.PostFormValue("password")
	//如果输入的密码加密后等于数据库查询出来的密码，则将用户存入session
	if user.Password == moudels.Encrypt(password) {
		session, err := user.CreateSession()
		if err != nil {
			fmt.Println("connot create session")
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

//GET /logout
//用户退出
func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != http.ErrNoCookie {
		fmt.Println("faild to get cookie")
		session := moudels.Session{UUID: cookie.Value}
		session.DeleteByUUID()
	}
	http.Redirect(w, r, "/", 302)

}
