package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	//解析模板
	tmpl, err := template.ParseFiles("F:\\goland\\go_project\\go_web\\websrc\\web_04\\hello.tmpl")
	if err != nil {
		fmt.Println("create tmplate failed,err:", err)
		return
	}
	//渲染模板
	tmpl.Execute(w, "沙河")
}
func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server is failed", err)
		return
	}
}
