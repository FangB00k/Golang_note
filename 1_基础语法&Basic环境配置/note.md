# 基础语法笔记
> 本文配环境 简略，请萌新自行百度配置环境，相应文章很多 本文只是一笔略过
><br> 配置环境文1：https://www.jb51.net/article/211334.htm
## 语言基础介绍
语言类型：
 - 强类型
 - 编译型语言
 - 不需要依赖，编译时自动打包进去


开发过的项目
- Docker
- 后台服务器

编译不同平台(通过环境变量控制编译)：
 - mac  GOOS=darwin
 - linux  GOOS=linux
 - windows  GOOS=windows

不同架构通过 GOARCH环境变量


go语言编译指令

```shell
 go build -o file.exe xx.go xx2.go
```
```shell
go build *.go
```
go语言直接运行
```shell
go run *.go
```
go install 文件
- [ ] http://c.biancheng.net/view/122.html
> go env 查看当前配置<br>
> 用的时候直接通过终端输入例如 GOODS =  xxx...
## 0x0 基础工程管理
  基础工程源码文档 ：GOPATH
   - src
     - 存放源码
   - bin[GOBIN<==环境变量]
     - 编译之后的程序存放位置
   - pkg
     - 包文件
### 0x00 环境变量配置
GOROOT
- 安装目录
   
GOPATH
- 工作空间，用来存放包目录

> 稍后更新
## 0x1 基础语法Start-HelloWolrd
> 如果之前学过C语言则可以快速通过对比语法区别入门 <br>
> C语言与go语言区别对比  https://hyperpolyglot.org/c

包含标准库文件
```go
import "fmt"
```
> 标准包名fmt，类似于C语言的 stdio.h,C+的iostream 包含了基本的输入输出

第一个HelloWorld

```go

package main
import "fmt"
func main(){
    fmt.Printf("Hello World!\n")
}
```
编译
> go build 文件名<br>
> ./文件名

<mark>对应 HelloWorld.go</mark>
<br><br>
一些小特性
- go语言不需要;号结尾！！ 
- package相当于 命名空间
    - 一般跟随文件夹名字
- 函数的花括号必须与函数同行


# 0x2 基础语法

## 0x21变量 var
  定义
  > var 变量名称 变量类型 <br> var 变量名称 变量类型 = 变量值
  赋值
  > 变量名 = 变量值

  一种额外的特性定义方法根据变量值自动定义变量的类型：
  > 变量名:=变量值


  <mark>对应 21var.go</mark>
  ```go
  package main
  import "fmt"
  func main()  {
    var var1 int = 1
    var var2 string = "哈哈"
    var3:="6666"
    fmt.Printf("var1:%d var2:%s var3:%s",var1,var2,var3)

  }
  ```

  赋值 
  - 同时赋值
    > 变量1,变量2 = 10,20
 
  变量类型
  - int
    - int8
    - int16
    - int32
    - int64
    - uint8
    - uint64
    - uint32
  - String
  - ture
  - float
    - float32
    - float64

  自增
  - 变量++
  - 变量--
  - 自增语法必须单独一行！！
  
字符串
- 简单赋值
  - > name:="abc"
- 多行赋值(等于python """)
```go
messg:=`  aaaa    bbbb 
   ccc ddd`
```
- 长度
  - len(字符串) 返回int 类型
- 拼接
  - a+b


变量类型转换
- 转换为浮点型 flaot64(变量)
- 转换为整型 uint(变量)






## 0x22指针 [->] [<-]
- [ ] 需要日后展开，先了解咋用
> go语言在使用指针的时候，会使用内部的垃圾回收机制，不需要手动释放内存。 
> go可以返回栈上的指针，变量分配到堆上
```go

package main
import "fmt"
func main(){
  var str string = "Hello"
  pointer:=&str;
  fmt.Printf("%s \n address:%0x",*pointer,pointer)

}

```
新建内存指针

c语言
```
int *pointer = (int *)malloc(sizefo(int))
```
go语言
```go
package main
import "fmt"
func main(){
	var star *string = new(string)
	*star = "star star"
	fmt.Println(*star)
}
```
go 语言
```go
Pointer :=new(string)
*Pointer = "Dark"
```

这tm的是啥狗屁🐕👇
<mark>可以返回栈上指针，编译器会在编译程序的时候，会自动判断这段代码，将city变量分配在堵上，内存逃逸</mark>

如何判断空指针
- go: nil 
- c: null

不支持语法
- 不支持地址加减
- 不支持三目运算符
- 只有false才能代码逻辑假，数字0和nil不能


## 0x23 数组 array
数组分类
- 定长数组
- 变长数组
### 0x231 定长数组
c语言
```c
int num[5]={1,2,3,4}
```

go语言
```go
num:=[5]int{1,2,3,4}
```

```go
var num = [10]int{1,2,3,4}
```

```go
var num [10]int = [10]int{1,2,3,4}
```
遍历
```go
for i:=0;i<len(数组);i++{
   fmt.Println("i:",i,",j:",数组[i])
}
```

```go
for key,value :=range nums{
  fmt.Println(key,value)
}
// 不会变因为 这个 key和value都是新的变量
```
```go
//想忽略一个值用_,两个都忽略用=而不是用：=
for _,value :=range nums{
  fmt.Println(value)
}
```


### 0x232 变长数组
 变长数组：顾名思义，就是可以动态的改变长度

定义
 ```go
   names:=[]string{"food","pen","book"}
 ```
追加长度
```go
names = append(names,"sword")
// 重新赋值，重新更新
```

容量的概念
- 为了避免多次申请内存，如果是在当前容量满状态追加，一次字节申请一倍的空间。。
- 查看容量代码：cap()
```go
var strys = []string{"xy","xz"}
fmt.Println("star1-容量",cap(strys))
strys = append(strys,"zy")
fmt.Println("star2-容量",cap(strys))
```
```markdown
> 输出
star1-容量 2
star2-容量 4
```

### 0x234 数组切片[变长]

切片
- 切片是左闭右开
- 切片修改后原数组也会变
- 如果是从0元素开始截取，冒号左边的数组可以省略 [0:5],[:5]
- 截取全部[:]
- 截取到尾部[开始坐标:]
- 取子串 "HaHaHa123"[0:2] = "Ha"
- 切片前 必须要有空切片


什么是空切片？
>“ 空切片是有分配内存但底层指向的是一个空数组 ”
```go
s1 := []int{}   //1.空切片，没有任何元素
s2 := make( []int, 0)  //2.make 切片，没有任何元素
```
空切片创建函数 make
```go
str2:=make([]string,10,20)
// 长度10，容量20，明确切片容量
```
浅拷贝 和 深拷贝
- 浅拷贝：共享地址
- 深拷贝：copy

浅拷贝切片代码
```go
package main
import "fmt"
func main(){
	location:=[18]string{"shanghai","shandong","shanxi","henan","hebei"}
	// fmt.Printf(location[1])
	// // 基于location创建一个新的数组
	// newb:=[]string{}
	// newb[0] = location[0]
	// newb[1] = location[1]
	// newb[2] = location[2]
	// 基于切片,创建一个数组
	newb2:=location[0:3]
    fmt.Println(newb2)
}

// func outarray(var star){
//    for var i:=1;i<len(star);i++{
// 	 Println(star[i])
//    }
// }
```
深拷贝：

- 通过cp
```go
package main
import "fmt"
func main(){
	newp:=[]string{"a","b","c","d","e","F"}
	newp2:=make([]string, 5,20)
	copy(newp2,newp[:3])

	for i:=0;i<len(newp2);i++{
		fmt.Println(newp2[i])

	}

}
```
## 0x24 map字典
> map[int]string 等价于 map[int] = string
```go
	var shopename map[int]string
  shopename[12]= "123";
	fmt.Printf(shopename[12])
```
1.定义一个map，但是map不能直接使用
- panic: assignment to entry in nil map

2.需要直接使用make地址进行分配空间

```go
package main
import "fmt"
func main(){
	var shopename map[int]string
	// 定义一个map，但是map不能直接使用
	// panic: assignment to entry in nil map

	shopename = make(map[int]string,10)
	shopename[12]= "123";
	fmt.Printf(shopename[12])
}
```

map的遍历
```go
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
}
```

判断是否存在

```go
	fmt.Println(shopename[100] == "")
  key,value:=shopename[100]
```

删除map函数

```go
delete(map函数,keyword)
```

## 0x25 Function 函数
> 可以返回多个返回值
- 返回单个值
```go
package main
import "fmt"
func main(){
	fmt.Println(aaddb(1,5))
}

func aaddb(a int,b int) int{
	return a+b
}
```
- 返回多个值


```go
package main
import "fmt"
func main(){
	fmt.Println(aaddb(1,5))
	fmt.Println(a_add_sub_b(1,5))

}

func a_add_sub_b(a int,b int)(int,int){
	return a+b,a-b
}

func aaddb(a int,b int) int{
	return a+b
}

```

- 返回多个值省略其中某个值

```go
package main
import "fmt"
func main(){
	fmt.Println(aaddb(1,5))
	a,_ := a_add_sub_b(1,5)
	fmt.Println(a)

}

func a_add_sub_b(a int,b int)(int,int){
	return a+b,a-b
}

func aaddb(a int,b int) int{
	return a+b
}

```
> 定义多个相同类型变量
> ```go
> var i,k,j int
>```

- 当返回值类型有了名称 就直接写return
```go
func a_sub_b(a int,b int) (result int){
	result = a - b
	return
}
```
## 0x26 内存逃逸
- 原本从栈上的内容逃跑到堆上
- 返回？◀

## 0X27 import包导入
- 默认调用包名 需要使用 **包名.函数**
- 如果一个包名想要对外提供访问权限一定要首字母大写
   - 首字母小写相当于 Private 只有相同包名才可以使用
   - 首字母大写相当于 Public
```go
 import(
  name "path/paths"  //name是给包命名
  . "path/paths" //直接引用
  "fmt"
 )

```


