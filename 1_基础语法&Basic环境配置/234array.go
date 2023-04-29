package main
import "fmt"
func main(){
	location:=[18]string{"shanghai","shandong","shanxi","henan","hebei"}
	// newb:=[]string{}
	// newb[0] = location[0]
	// newb[1] = location[1]
	// newb[2] = location[2]
	newb2:=location[0:3]
    fmt.Println(newb2)

}

// func outarray(var star){
//    for var i:=1;i<len(star);i++{
// 	 Println(star[i])
//    }
// }