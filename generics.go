package main
import "fmt"
func printslice[T int|string](items []T){
	for _,item:=range items{
		fmt.Println(item)
	}
}
type stack [T any]struct{
	elements []T
}
func main(){
	nums:=[]int{1,2,3}
	strs:=[]string{"min","haj"}
	printslice(nums)
	printslice(strs)
	mystack:=stack[string]{
		elements:[]string{"min","haj"},
	}
	fmt.Println(mystack.elements)
}