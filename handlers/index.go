package handlers

import (
	"io"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	log.Println("listen by address", r.RemoteAddr)
	//w.WriteHeader(http.StatusOK)
	io.WriteString(w, "hello index")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world")
}
