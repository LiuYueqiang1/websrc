package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	k := func(name string) (string, error) {
		return name + "年轻又帅气", nil
	}
	//创建模板
	t := template.New("f.tmpl")
	t.Funcs(template.FuncMap{
		"kua": k,
	})
	//解析模板
	tmpl, err := t.ParseFiles("F:\\goland\\go_project\\go_web\\websrc\\web_06\\f.tmpl")
	if err != nil {
		fmt.Println("create tmplate failed,err:", err)
		return
	}
	//渲染模板
	t.Execute(w, "帅气")    //渲染一次  这两个都可以
	tmpl.Execute(w, "沙河") //渲染一次
}
func demo1(w http.ResponseWriter, r *http.Request) {
	//*******加载文件的先后顺序十分重要*******
	t, err := template.ParseFiles("F:\\goland\\go_project\\go_web\\websrc\\web_06\\t.tmpl", "F:\\goland\\go_project\\go_web\\websrc\\web_06\\u1.tmpl")
	if err != nil {
		fmt.Printf("parse tmplate failed,err:%v\n", err)
		return
	}
	name := "小王子"
	t.Execute(w, name)
}
func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/tmplDemo", demo1)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:%v", err)
		return
	}
}
