package main
import "fmt"
func main(){
	i:=1
	//while loop
	for i<=5{
		fmt.Println(i)
		break
		i=i+1
	}
	//for loop
	for i:=0;i<=3;i++{
		if i==2{
			continue
		}
		fmt.Println(i)
		
	}
}