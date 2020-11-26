package handlers

import (
	"log"
	"main/moudels"
	"net/http"
)

//论坛首页路由处理
func Index(w http.ResponseWriter, r *http.Request) {
	log.Println("listen by address", r.RemoteAddr)
	//w.WriteHeader(http.StatusOK)
	//io.WriteString(w, "hello index")
	//files := []string{"views/index.html", "views/layout.html", "views/navbar.html"}
	//templates := template.Must(template.ParseFiles(files...))
	threads, err := moudels.Threads()
	if err == nil {
		_, err := session(w, r)
		if err == nil {
			//log.Fatal(err.Error())
			generateHtml(w, threads, "layout", "index", "auth.navbar")
		} else {
			generateHtml(w, threads, "layout", "index", "navbar")
		}
	}

}
