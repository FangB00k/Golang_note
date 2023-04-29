package main
import "fmt"
func main(){
  var str string = "Hello"
  pointer:=&str;
  potinter2:=new(string)
  *potinter2 ="star⭐"
  fmt.Println(*potinter2)
  fmt.Printf("%s \n address:%0x",*pointer,pointer)

}