package main
import "fmt"
func main(){
fmt.Println("欢迎大家来玩耍哦！")
for i:=0;i<10;i++{
	fmt.Println("循环！")
	goto label1
}

label1:
  fmt.Println("我跳出来了！")
}