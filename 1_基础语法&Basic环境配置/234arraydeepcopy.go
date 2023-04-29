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