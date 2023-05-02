package main
import "fmt"
// import "time"

func main(){
	numChan := make(chan int) 
	// 如果make(chan int) 我们称为这个是无缓冲的
	// 如果make(chan int,10)这个我们称为缓冲的

 	// 装数字的管道，使用管道的时候一定make，同map一样，否则 是 nil
	// 儿子读数据
	go func(){
		for i:=0;i <50;i++{
			data:=<- numChan
			fmt.Println("data:",data)
		}
	}()

	// 创建 两个go程 父亲写数据,儿子读数据

	for i:=0;i<50;i++{
		// 向管道写入数据
		numChan <- i
		fmt.Println("这是主进程写入了:",i)
	}




}