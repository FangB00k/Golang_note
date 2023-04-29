package main
import "fmt"
func main(){
   //=====定长 
   var num = [10]int{1,2,3,4}
   nums:=[10]string{"sword","beer"}
   var numx[10]int =[10]int{1,2,3,4}

   for i:=0;i<10;i++{
	fmt.Println(num[i]);
   }

   for _,value:= range nums{
	fmt.Println(value)
   }

   for _,value:=range numx{
	 fmt.Println(value)
   }
   //========变长

   var strx = []string{"x","y"}
   stry :=[]string{"aa","bb"}
   strx = append(strx,"z")
   for i:=0;i < len(strx);i++{
	 fmt.Println(strx[i])
   }
   for i:=0;i < len(stry);i++{
	fmt.Println(stry[i])
  }
// ========= 容量

var strys = []string{"xy","xz"}
fmt.Println("star1-容量",cap(strys))
strys = append(strys,"zy")
fmt.Println("star2-容量",cap(strys))


}