package handlers

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

//论坛首页路由处理
func Index(w http.ResponseWriter, r *http.Request) {
	log.Println("listen by address", r.RemoteAddr)
	//w.WriteHeader(http.StatusOK)
	//io.WriteString(w, "hello index")
	files := []string{"views/index.html", "views/layout.html", "views/navbar.html"}
	templates := template.Must(template.ParseFiles(files...))
	//threads, err := moudels.Threads()
	//if err!=nil {
	//	log.Fatal(err.Error())
	//}
	//if err==nil {
	//	templates.ExecuteTemplate(w,"layout",threads)
	//}
	templates.ExecuteTemplate(w, "layout", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world")
}
