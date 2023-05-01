package main
import "fmt"
type Integar int

type student struct{
	name string
	age int 
	sorce int
}
func main(){
	var i Integar
	i = 10
	fmt.Println(i)
	var x student
	x.name = "哈哈哈"
	x.age = 10
	x.sorce = 100
	fmt.Println(x)
	var y student = student{"小明",10,50}
    fmt.Println(y)
}
