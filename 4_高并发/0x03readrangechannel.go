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
	   // 🍅
	   fmt.Println("数据全部写完毕，准备关闭管道！")
	   // 🍅
	   close(numsChan)

	}()

	// 遍历轨道的时候只返回一个值 	
	// for range 是不值得管道是否已经写完了，所以会一直在这里等着
	 // 解决办法 添加关闭 🍅
	for v := range numsChan{
		fmt.Println("读取数据:",v)
	}
}