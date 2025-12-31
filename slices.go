package main
import "fmt"
func main(){
	var nums[]int
	fmt.Println(len(nums))

	var num=make([]int,3,50)
	num=append(num,6)
	fmt.Println(len(num))
	fmt.Println(num)
}