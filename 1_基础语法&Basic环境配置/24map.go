package main
import "fmt"
func main(){
	var shopename map[int]string
	// 定义一个map，但是map不能直接使用
	// panic: assignment to entry in nil map

	shopename = make(map[int]string,10)
	shopename[12]= "123";
	// fmt.Printf(shopename[12])
	shopename[78]="789"
	for key,value:= range shopename{
		fmt.Println(key,value)
	}

	fmt.Println(shopename[100] == "")
}