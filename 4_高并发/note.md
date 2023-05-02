# 0x00 概念
并发： 同时运行多个任务 然后通过时钟切片不断切换，实现感觉运行了很多
并行： 是CPU利用多核 同时执行

- C语言并发 使用了多线程 ， 进程
- Go语言并发，使用了`Go程`？ go语言原生支持的

相比:
- go语言里不是线程 占用资源小于线程，一个go程 需要 4K~5K内存，一个线程可以启动大量的go程序

使用:
 - 只需要在目标函数前面加上go关键字就够了


> 1.创建例子

- 主线程（main）
- 子线程

```go

package main
import (
	"fmt"
	"time"
)

func display(){
	//子程序
	count:=1
	for{
		fmt.Println("===> this 子go程序:",count)
		count++
		time.Sleep(1500)
	}

}

func main(){
  // 主线程

  //启动子程序
  go display()

  // 主go程序
  count:=1
  for{
	  fmt.Println("===> this 主go程序:",count)
	  count++
	  time.Sleep(1*time.Second)
  }
}

```
或者直接把子函数写成匿名函数
```go

package main
import (
	"fmt"
	"time"
)

func display(){
	//子程序
	count:=1
	for{
		fmt.Println("===> this 子go程序:",count)
		count++
		time.Sleep(1500)
	}

}

func main(){
  // 主线程

  //启动子程序
  go func ()  {
	count:=1
	for{
		fmt.Println("===> this 子go程序:",count)
		count++
		time.Sleep(1500)
	}

  }()

  // 主go程序
  count:=1
  for{
	  fmt.Println("this 主go程序:",count)
	  count++
	  time.Sleep(1*time.Second)
  }
}

```

启动多个会公平竞争

```go

package main
import (
	"fmt"
	"time"
)

func display(num int){
	//子程序
	count:=1
	for{
		fmt.Println("===> this",num," 子go程序:",count)
		count++
		time.Sleep(1500)
	}

}

func main(){
  // 主线程
   
  // 启动子进程
  for i:=1;i<=3;i++{
	go display(i)
  }
  

  // 主go程序
  count:=1
  for{
	  fmt.Println("this 主go程序:",count)
	  count++
	  time.Sleep(1*time.Second)
  }
}

```
# 0x01 return 和 go exit 区别
- 使用GoEXIT提前退出当前go程
  - runtime.Goexit()
- return 函数返回
- exit 推出进程

# 0x02 多进程（go程序) 通信
> 多线程的时候 c语言使用互斥量，上锁保持资源同步，避免资源竞争问题
><br> go语言使用管道
> <br> A往通道里面写数据 B从管道里面读数据，go自动帮我们你做好了数据同步

- 无缓冲管道

1. 创建管道
2. 读入数据

```go

package main
import "fmt"
// import "time"

func main(){
	numChan := make(chan int) 
	// 如果make(chan int) 我们称为这个是无缓冲的
	// 如果make(chan int,10)这个我们称为缓冲的

 	// 装数字的管道，使用管道的时候一定make，同map一样，否则 是 nil
	// 儿子读数据
	go func(){
		for i:=0;i <50;i++{
			data:=<- numChan
			fmt.Println("data:",data)
		}
	}()

	// 创建 两个go程 父亲写数据,儿子读数据

	for i:=0;i<50;i++{
		// 向管道写入数据
		numChan <- i
		fmt.Println("这是主进程写入了:",i)
	}


```

- 读写问题
```go

package main
func main(){
	numsChan := make(chan int,10)
	// 1. 当缓冲区满的时候，写阻塞，当被读取后，再恢复写入
	// 2. 当缓冲区读取完毕后，读阻塞
	// 3.如果管道没有使用make分配空间,那么管道默认是nil的 读取和写入都阻塞
    // 4.对于一个管道，读与写的次数，必须对等
	//  - 如果阻塞在写上 那么会推出
	//  - 如果阻塞在子程序上会 内存溢出
	var names chan string //默认是nil的
	names <- "hello" //由于names是nil的，写操作会阻塞在这里
	fmt.Println("names:",<-names)



}

```

# 0x03 range读取管道

- 遍历轨道的时候只返回一个值 	
- for range 是不值得管道是否已经写完了，所以会一直在这里等着
- 解决办法 添加关闭 ,关闭管道🍅
```go
package main
import "fmt"
func main(){
	numsChan :=make(chan int, 10)
	// 写
	go func(){
	   for i:=0;i<50;i++{
		 numsChan <- i
		 fmt.Println("写入数据:",i)
	   }	
	   // 🍅
	   fmt.Println("数据全部写完毕，准备关闭管道！")
	   // 🍅
	   close(numsChan)

	}()

	// 遍历轨道的时候只返回一个值 	
	// for range 是不值得管道是否已经写完了，所以会一直在这里等着
	 // 解决办法 添加关闭 🍅
	for v := range numsChan{
		fmt.Println("读取数据:",v)
	}
}
```

1. 向一个已经 关闭的通道 close 读取数据 返回零值 
2. 向一个已经close的管道写入数据,会崩溃
3. 关闭管道的动作写在写端


# 0x04 判断管道是否已经关闭
- 判断管道

`失败我靠...🍅`
```go

package main

import "fmt"
func main(){
	numsChan :=make(chan int, 10)
	// 写
	go func(){
	   for i:=0;i<50;i++{
		 numsChan <- i
		 fmt.Println("写入数据:",i)
	   }	
	   fmt.Println("数据全部写完毕，准备关闭管道！")
	   close(numsChan)

	}()

	for{
		v,ok := <- numsChan
		if !ok{
			fmt.Println("管道已经关闭了")
			break
		}
		fmt.Println("读取的数据:",v)
	}
}
```


# 0x05 单项管道

numChan:= make(char int,10) => 双向通道即可以读又可以写

单项通道 
 - 单项读通道 var numChanRead <- chan int
 - 单项写通道 var numChanRead chan<-int
单向通道顾名思义，只能写或者只能读


# 0x06 Select

当程序中有多个 channel协同工作,被触发了就会被监听到

```go
package main
import "fmt"
import "time"
func main(){
	chan1:=make(chan int, 10)
	chan2:= make(chan int,10)

	// 启动一个go程监听两个 

	go func(){
		for{
			select{
			case data1:= <-chan1:
				fmt.Println("chan1读取了数据嘞<",data1)
			case data2:=<-chan2:
				fmt.Println("chan2读取了数据嘞<",data2)
			
			}
		}
	}()


	// 启动 chann1写入数据
	go func(){
		for i:=0;i<10;i++{
			chan1 <-i
			fmt.Println("chan1写入数据")
			time.Sleep(1000)
		}
	}()

	
	// 启动 chann2写入数据
	go func(){
		for i:=0;i<10;i++{
			chan2 <-i
			fmt.Println("<<chan2写入数据")

			time.Sleep(1000)
		}
	}()

}
```