package main
import "fmt"
// 定义一个接口 
type IAttack interface{
   // 接口函数可以有多个 到那时只能有函数原型 不可以有实现
   Attack()
}
type Person struct{
	name string
	level int
}



func (this *Person) Attack(){
	fmt.Println(this.name + "send Attact")
}

func main(){
  human:=Person{name:"fang",level:100}
  human.Attack()

  // 创建一个接口
  var player IAttack
  player = &Person{name:"fang",level:10} // 只要实现了 IAttack的类就可以通过地址赋值给接口
  // 接口只能接受地址
  player.Attack()
}