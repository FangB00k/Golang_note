package main
import "fmt"
type Person struct{
	name string
	age int 
	gender string
	scorce float64
}

func (p Person) Eat(){
	fmt.Println("Eat good fooding....")
}

func (this *Person) Run(){
	fmt.Println(this.name+" runing runing!")
	return
}

func main(){
	xiaoming:=Person{name:"xiaoming",age:10,gender:"two",scorce:114.4,}
	xiaoming.Eat()
	xiaoming.Run()
}