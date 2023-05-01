package main
import "fmt"
func main(){
	var star *string = new(string)
	*star = "star star"
	fmt.Println(*star)
}