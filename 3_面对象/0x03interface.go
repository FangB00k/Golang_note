package main
import "fmt"
func main(){
	var i interface{} // 定义1个接口
	school:=[]string{"dalei","jojo"}
	i = school
	fmt.Println(i)

	// 创建一个只有三个接口的切片
	array:=make([]interface{}, 3)
	array[0] = 1
	array[1] = "Hello WOlrd"
	array[2] = 11.4
	for _,value := range array{
		v:=value.(type)
		switch  v{
			case int:	
			  	fmt.Println("当前类为int")
			case string:
				fmt.Println("当前为string")
			case float64:
				fmt.Println("当前为float")
			default:
				fmt.Println("非法数据")
		}
	}
}