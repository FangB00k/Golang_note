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