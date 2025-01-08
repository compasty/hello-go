package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func readBaidu() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	// 需要关闭连接
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func examplePost() {
	body := "{\"action\":20}"
	res, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer([]byte(body)))
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	defer res.Body.Close()

	content, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	fmt.Println(string(content))
}

func SayHello(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	w.Write([]byte("Hello," + name + "!"))
}

func main() {
	readBaidu()
	examplePost()
	// 建立服务器
	http.HandleFunc("/", SayHello)
	http.ListenAndServe(":8123", nil)
}
