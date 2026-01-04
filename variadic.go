package main
import "fmt"
func sum(nums ...int )int{
	total:=0
	for _,i := range nums{
		total=total+i
	}
	return total
}
func main() {
	fmt.Println(sum(1,2,3,4,5,6))
}
