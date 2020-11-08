package main

import (
	"log"
	. "main/router"
	"net/http"
)

func main() {
	startWebServer("8080")
}

func startWebServer(port string) {
	r := NewRouter()
	//处理静态文件
	assets := http.FileServer(http.Dir("public"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", assets))
	http.Handle("/", r)

	log.Println("satrt hhtp service at port:", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Println("start http server error", err.Error())
	}
}
