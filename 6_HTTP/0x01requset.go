package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := http.Client{}
	resp, err := client.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("net error:", err)
		return
	}

	date := resp.Header.Get("Date")
	server := resp.Header.Get("Server")
	ct := resp.Header.Get("Content-Type")
	fmt.Println(date)
	fmt.Println(server)
	fmt.Println(ct)
	body := resp.Request.Body
	fmt.Println(body)
	url := resp.Request.URL
	fmt.Println(url)
	//读取网页
	res, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(res))

}
