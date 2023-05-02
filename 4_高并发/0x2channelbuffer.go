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