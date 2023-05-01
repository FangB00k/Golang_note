package main
import "fmt"
func main(){
	const(
		ONE = iota+1
		TOW = 99
		XX = iota
		THREE
		FOUR
	)
	i:= THREE
	fmt.Println(i)
	fmt.Println(TOW)
}