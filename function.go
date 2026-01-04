package main
import "fmt"

func getlang()(string,string,string){

	return "go","python","near"
}
func main(){
	s,d,f:=getlang()
	fmt.Println(s,d,f)

}