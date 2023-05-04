package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.liwenzhou.com/")
	if err != nil {
		fmt.Println("init http failed!")
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("print http failed!")
		return
	}
	fmt.Print(string(body))
}
