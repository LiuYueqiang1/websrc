package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	//创建模板
	//解析模板
	t, err := template.New("index.tmpl").Delims("{[", "]}").ParseFiles("./websrc/web_08/index.tmpl")
	if err != nil {
		fmt.Println("parse template failed,err:", err)
		return
	}
	//渲染模板
	t.Execute(w, "马老师")
}
func xss(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(str string) template.HTML { //函数名：函数
			return template.HTML(str)
		},
	}).
		ParseFiles("./websrc/web_08/xss.tmpl")
	if err != nil {
		fmt.Println("parse template failed,err:", err)
		return
	}
	str1 := "<script>alert(123)</script>"
	srt2 := "<a href='http://liwenzhou.com'>liwenzhou的博客</a>"
	t.Execute(w, map[string]string{
		"s1": str1,
		"s2": srt2,
	})
}
func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/xss", xss)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server is failed", err)
		return
	}
}
