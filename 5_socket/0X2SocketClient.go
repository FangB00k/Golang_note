package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8848")
	if err != nil {
		fmt.Println("链接失败:", err)
		return
	}
	fmt.Println("建立链接成功！")

	sendData := "Hello World"
	conn.Write([]byte(sendData))
	if err != nil {
		fmt.Println("发送书失败:", err)
		return
	}
	fmt.Println("发送成功")

	// 接收返回数据
	buf := make([]byte, 1024)
	cnt, err := conn.Read(buf)
	if err != nil {
		fmt.Println("接收的时候发现错误?:", err)
		return
	}

	fmt.Println("接收的数据为:", cnt, string(buf))

}
