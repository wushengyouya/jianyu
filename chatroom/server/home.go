package server

import (
	"fmt"
	"net/http"
	"text/template"
)

func homeHandleFunc(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles(rootDir + "/template/home.html")
	if err != nil {
		fmt.Fprint(w, "模版解析错误")
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		fmt.Print(w, "模版执行错误")
		return
	}
}
