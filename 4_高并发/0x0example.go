package main
import (
	"fmt"
	"time"
)

func display(){
	//子程序
	count:=1
	for{
		fmt.Println("===> this 子go程序:",count)
		count++
		time.Sleep(1500)
	}

}

func main(){
  // 主线程

  //启动子程序
  go func ()  {
	count:=1
	for{
		fmt.Println("===> this 子go程序:",count)
		count++
		time.Sleep(1500)
	}

  }()

  // 主go程序
  count:=1
  for{
	  fmt.Println("this 主go程序:",count)
	  count++
	  time.Sleep(1*time.Second)
  }
}