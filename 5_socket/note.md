# 0x00 socket的过程

(●ˇ∀ˇ●)我是客户端：

> 1. 指明我要和谁通话
> 2. 写如数据
> 3. 等待响应 
> 4. 关闭链接

```go
package main
import "net"
import "fmt"
func main(){
	conn,err := net.Dial("tcp","127.0.0.1:8848")
	if err != nil {
		fmt.Println("链接失败:",err)
		return
	}
	fmt.Println("建立链接成功！")
	
	sendData:="Hello World"
	conn.Write([]byte(sendData))
	if err != nil {
		fmt.Println("发送书失败:",err)
		return
	}
	fmt.Println("发送成功")


	// 接收返回数据
	buf:=make([]byte,1024)
	cnt,err:=conn.Read(buf)
	if err != nil {
		fmt.Println("接收的时候发现错误?:",err)
		return
	}

	fmt.Println("接收的数据为:",cnt,string(buf))

}
```

(｡･∀･)ﾉﾞ我是服务端：

> 1. 监听谁要和我通话
> 2. 收到后读出来
> 3. 返回响应
> 4. 关闭链接🔗

```go
package main
import "net"
import "fmt"
func main(){
	// 创造监听
	ip:="127.0.0.1"
	port:=8848
	address:=fmt.Sprintf("%s:%d",ip,port)
	listener,err := net.Listen("tcp",address)
	if err != nil {
		fmt.Println("netserver_error:",err)
		return
	}
	conn,err  := listener.Accept()
	if err != nil {
		fmt.Println("netserver_error_accept:",err)
		return
	}
	fmt.Println("链接建立成功！")
	//创建一个容器用于接收读取到的数据
	buf:=make([]byte, 1024)
	cnt,err:=conn.Read(buf)

	if err != nil {
		fmt.Println("netserver_error_accept1:",err)
		return
	}

	fmt.Println("Client ==> Server,len",cnt,"data:",string(buf))
	//往回写
	uperData:= "OK"
	cnt,err = conn.Write([]byte(uperData))

	// 关闭链接
	conn.Close()


}
```

