package main
import "fmt"
type Person struct{
	name string
	age int 
}
func (this *Person)eat(){
	fmt.Println(this.name + " love eat food！")
}
// 嵌套
type Student struct{
	person Person
	source int
}

// 继承
type Teacher struct{
	Person
	school string

}
func main(){
  xiaoming:= Student{person:Person{"hahah",10},source:50}
  fmt.Println(xiaoming.person.name)
  xiaoming.person.eat()
  xiaowang:= Teacher{} // {name:"xiaowang",age:25,school:"diqiu"}
  xiaowang = Teacher{Person:Person{name:"xiaowang"}}
  xiaowang.eat()

}