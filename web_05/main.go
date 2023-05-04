package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Student struct {
	Name    string
	Age     int
	Address string
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//解析模板
	tmpl, err := template.ParseFiles("F:\\goland\\go_project\\go_web\\websrc\\web_05\\hello.tmpl")
	if err != nil {
		fmt.Println("create tmplate failed,err:", err)
		return
	}
	student := Student{
		Name:    "沙河",
		Age:     18,
		Address: "beijing",
	}
	stu := map[string]string{
		"name": "ss",
		"addr": "newyork",
	}
	hobbyLise := []string{
		"闪电五连鞭",
		"偷袭",
		"大意了啊",
	}
	//渲染模板
	tmpl.Execute(w, map[string]interface{}{
		"u1":    student,
		"s1":    stu,
		"hobby": hobbyLise,
	})
}
func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server is failed", err)
		return
	}
}
