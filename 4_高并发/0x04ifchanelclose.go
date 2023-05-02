package main

import "fmt"
func main(){
	numsChan :=make(chan int, 10)
	// 写
	go func(){
	   for i:=0;i<50;i++{
		 numsChan <- i
		 fmt.Println("写入数据:",i)
	   }	
	   fmt.Println("数据全部写完毕，准备关闭管道！")
	   close(numsChan)

	}()

	for{
		v,ok := <- numsChan
		if !ok{
			fmt.Println("管道已经关闭了")
			break
		}
		fmt.Println("读取的数据:",v)
	}
}