package main
import "fmt"
import "time"
func main(){
	chan1:=make(chan int, 10)
	chan2:= make(chan int,10)

	// 启动一个go程监听两个 

	go func(){
		for{
			select{
			case data1:= <-chan1:
				fmt.Println("chan1读取了数据嘞<",data1)
			case data2:=<-chan2:
				fmt.Println("chan2读取了数据嘞<",data2)
			
			}
		}
	}()


	// 启动 chann1写入数据
	go func(){
		for i:=0;i<10;i++{
			chan1 <-i
			fmt.Println("chan1写入数据")
			time.Sleep(1000)
		}
	}()

	
	// 启动 chann2写入数据
	go func(){
		for i:=0;i<10;i++{
			chan2 <-i
			fmt.Println("<<chan2写入数据")

			time.Sleep(1000)
		}
	}()

}