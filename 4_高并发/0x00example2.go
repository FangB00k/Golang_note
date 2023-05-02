package main
import (
	"fmt"
	"time"
)

func display(num int){
	//子程序
	count:=1
	for{
		fmt.Println("===> this",num," 子go程序:",count)
		count++
		time.Sleep(1500)
	}

}

func main(){
  // 主线程
   
  // 启动子进程
  for i:=1;i<=3;i++{
	go display(i)
  }
  

  // 主go程序
  count:=1
  for{
	  fmt.Println("this 主go程序:",count)
	  count++
	  time.Sleep(1*time.Second)
  }
}