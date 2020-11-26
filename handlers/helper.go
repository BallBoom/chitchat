package handlers

import (
	"errors"
	"fmt"
	"html/template"
	"main/moudels"
	"net/http"
)

// 通过cookie 用来判断用户是否已登陆
func session(w http.ResponseWriter, r *http.Request) (ses moudels.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		ses = moudels.Session{UUID: cookie.Value}
		if ok, _ := ses.Check(); !ok {
			err = errors.New("Invalid session!")
		}
	} else {
		fmt.Println("error", err.Error())
	}
	return
}

func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, f := range filenames {
		files = append(files, fmt.Sprintf("views/%s.html", f))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

// 获取HTML页面模板
func generateHtml(w http.ResponseWriter, datas interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/%s.html", file))
	}
	t := template.Must(template.ParseFiles(files...))
	t.ExecuteTemplate(w, "layout", datas)
}

func Version() string {
	return "1.0"
}
