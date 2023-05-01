# 0x1 基础语法结构
> C: argc, ** argv
> GO： os.Args 
```go
package main
import ("fmt"
"os")
func main(){
	cmds:= os.Args
	for key,cmd:= range cmds{
		fmt.Println("key:",key,"cmd:",cmd)
	}
}
```
## 0x12 Switch 语句
> 语法结构体
>  - switchu已经自动有break了不需要手动加了
>  - 向下穿透 加上 fallthrough
```go
switch exp[表达式]{
    case 值1:
      ....
    case 值2:
      ....
    case 值3:
      ....
    default:
      ....

}
```
## 0x13 GoTo&&Lable标签🏷
- Go to 相互跳转位置

> 打标签： 标签名:
> 跳转:
> - goto 标签
>   - 直接跳转走
> - continue 标签
>   - 保存状态然后跳转
> -  break 标签
>   - 直接跳出指定位置的循环
> 

```go
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
```


## 0x14 枚举类型
> go语言种没有枚举类型,批量定义常量const + iota(常量累加器)进行模拟

- 批量定义变量

```go

var number int
var boolth bool
var chars string
// 批量定义
var(
    number int 
    boolth bool
    char
)

```


- 批量定义常量+iota常量组计数器
> iota 换行才会加1 <br>
> 如果遇到 const iota会清零

```go
const(
    MONDAY = iota //0
    TUESDAY = iota //1
    WEDNESDAY // 2
    THURSDAY // 3
    FRIDAY // 4
    STATURDAY // 5 
    SUNDAY // 6 
)
```

> 直接赋值 不想从 0开始枚举

```go
const(
    JANU = iota+1 // 从1开始枚举
    FER // 2
    MAR // 3
    APRI // 4
)
```
  
## 0X15 Struct 结构体
> C语言: typedef 类型 类型名称 <br>
> Go语言: type 类型名称 类型
```go
package main
import "fmt"
type Integar int
func main(){
	var i Integar
	i = 10
	fmt.Println(i)
}
```
>  Go语言结构体

Go语言结构体 = type + struct 

举个例子 定义 学生结构体
```go
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

```

```go
student{"小明",
10,
50,}
```
如果是这种方式进行赋值 那么最后一行一定带一个逗号

## 0x16 init函数
- init函数会被引用(import)的时候自动调用
  
- 没有参数，没有返回值 ，可以写多个

- 如果只想调用包里面的init

```go
 import( 
    _ "包名"
 )
```


## 0x17 defer延迟
延迟用于修饰语句、函数确保这条语句可以在当前栈退出的时候执行
1. 解锁、关闭文件
2. 做资源清理工作
3. 在同一个函数中多次调用defer，执行时类似于栈的极值机制

go语言一般会将错误码最为最后一个参数返回
" defer 语句会将其后面跟随的语句进行延迟处理. 意思就是说 跟在defer后面的语言 将会在程序进行最后的return之后再执行 "

```go

lock.Lock()
a = "helo"
lock.Unlock()

````

```go

{
    locl.Lock()
    defer lock.UnLock()
    a = "hello"
}

{
    f1,_:= file.Open()
    defer f1.Close
}
```

匿名函数
- 如果一个函数没有文件名则是匿名函数
```go
def func(){
 匿名函数内容...
}()
```
## 0x18 if

```go
if expr {
    ..
}else if epxr{

}else{
    ...
}
```


## 0x19 循环

```go
for expr{
    ..
}

for 变量:=值;条件;变量++/--{

}

for 变量1,变量2:= range xxx{
    
}
```


