package main
import "fmt"
func main(){
	var hi =[3]int{2,5,6}
	mat :=[2][2]int{{2,2},{2,2}}
	hello :=[...]int{5,6,4,8}
	hi[1] =6
	fmt.Println(hi)
	fmt.Println(len(hello))
	fmt.Println(len(mat))
}