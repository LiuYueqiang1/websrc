package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadFile("websrc/gin_demo/hello.txt")
	fmt.Fprintln(w, string(b))
}
func main() {
	http.HandleFunc("/hello", sayHello)      //127.0.0.1:9090/hello 调用sayHello这个函数
	err := http.ListenAndServe(":9090", nil) //监听并且服务9090端口
	if err != nil {
		fmt.Printf("http server failed,err:%v\n", err)
		return
	}
}
