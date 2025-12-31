package main
func main(){
	i:=5
	switch i {
	case 1:
		fmt.Println(1)
case 2:
	fmt.Println(2)
case 3:
	fmt.Println(3)
case 5:
	fmt.Println(5)
default :
fmt.Println("no")
}
whoAmI := func(i interface{}) {

    switch i := i.(type) {

    case int:
        fmt.Println("its an integer")

    case string:
        fmt.Println("its a string")

    case bool:
        fmt.Println("its a boolean")

    default:
        fmt.Println("other")
    }
}
whoAmI(10)
whoAmI("hello")
whoAmI(true)
whoAmI(3.14)


}