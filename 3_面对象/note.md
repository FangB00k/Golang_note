# 0X01 类的创建
> 借助struct构建 对象

例子构造 Person对象
```go
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



func main(){
	xiaoming:=Person{name:"xiaoming",age:10,gender:"two",scorce:114.4,}
	xiaoming.Eat()
}
```

通过 **p Person** 来绑定方法
或者通过**this *Person**来绑定方法

```go
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
```


# 0x02 类的继承 和嵌套定义
嵌套：
 - 直接在里面声明 类型变量
继承 
 - 直接声明 一个 要继承的类型
 - 默认会创建一个默认同名的字段方便赋值
    - >   xiaowang = Teacher{Person:Person{name:"xiaowang"}}

关于作用域
 - 对于类和成员、方法 ，只有大写开头的才能在其他包中使用
 - 字母开头大写就是public  字母开头小写就是 private
```go

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
```

# 0x03 接口 interface
- Interface 不仅仅适用于处理多态 他可以接受任意的数据类型，有点类似void
- 空接口
举个例子
```
package main
import "fmt"
func main(){
	var i interface{} // 定义1个接口
	school:=[]string{"dalei","jojo"}
	i = school

	fmt.Println(i)
}
```
<mark>可以接收任意类型</mark>
判断类型
>值,是否是  =  接口.(类型)

常用方式
- 把interface当作一个函数参数 使用siwtch判断类型 根据不同类型作不同逻辑处理

例子

> 有bug啊啊啊啊啊，不过大概也是这样使用的
```go
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
```

# 0x04 多态
- 实现go多态需要实现接口
```go
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

```