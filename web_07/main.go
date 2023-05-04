package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./websrc/web_07/index.tmpl")
	if err != nil {
		fmt.Println("parse template failed,err:", err)
		return
	}
	msg := "闪电五连鞭"
	//渲染模板
	t.Execute(w, msg)
}
func home(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./websrc/web_07/home.tmpl")
	if err != nil {
		fmt.Println("parse template failed,err:", err)
		return
	}
	msg := "马老师"
	//渲染模板
	t.Execute(w, msg)
}
func index2(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./websrc/web_07/tmplates/base.tmpl", "./websrc/web_07/tmplates/index.tmpl")
	if err != nil {
		fmt.Println("parse template failed,err:", err)
		return
	}
	msg := "闪电五连鞭"
	//渲染模板
	t.ExecuteTemplate(w, "index.tmpl", msg)
}
func home2(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./websrc/web_07/tmplates/base.tmpl", "./websrc/web_07/tmplates/home.tmpl")
	if err != nil {
		fmt.Println("parse template failed,err:", err)
		return
	}
	msg := "马老师"
	//渲染模板
	t.ExecuteTemplate(w, "home.tmpl", msg)
}
func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server is failed", err)
		return
	}
}
