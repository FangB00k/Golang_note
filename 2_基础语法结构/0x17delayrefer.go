package main
import "fmt"
import "os"

func main()  {
  	fmt.Println("123")
	readFile("17.txt")
}

func readFile(filename string){
	f1,err := os.Open(filename)
	defer _=f1.close()
	if err!=nil{
		fmt.Println(filename,"打开失败==>",err)
		return
	}
	defer  fmt.Println("456")
	defer  fmt.Println("789")
	defer  fmt.Println("111213")

	buf:=make([]byte,1024)
	n,err := f1.Read(buf)
     
	fmt.Println("文件长度:",n)
	fmt.Println("读取文件内容:",string(buf))

	return


}