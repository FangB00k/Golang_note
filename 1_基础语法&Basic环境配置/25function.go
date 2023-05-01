package main
import "fmt"
func main(){
	// fmt.Println(aaddb(1,5))
	// a,_ := a_add_sub_b(1,5)
	// fmt.Println(a)
	fmt.Println(a_sub_b(5,1))

}

func a_add_sub_b(a int,b int)(int,int){
	return a+b,a-b
}

func aaddb(a int,b int) int{
	return a+b
}

func a_sub_b(a int,b int) (result int){
	result = a - b
	return
}