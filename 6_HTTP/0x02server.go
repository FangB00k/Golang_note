package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/hah", func(w http.ResponseWriter, r *http.Request) {
		//w 返回服务器信息
		//r 是客户端发来的请求数据
		fmt.Println("请求数据", r)
		// 写入数据
		_, _ = io.WriteString(w, "Hello World")
	})
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		fmt.Println("err", err)
	}

}
