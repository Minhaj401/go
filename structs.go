package main
import "fmt"
import "time"
type order struct {
	id string 
	amount float32
	status string
	cretatedat time.Time
}
//reciver type
func (o order) changestatus(status string){
	o.status=status
}
func (o order) neworder (id string ,amount float32,status string,cretatedat time.Time)*order{
	myorder:=order{
		id:id,
		amount:amount,
		status: status,
		cretatedat:time.now()
	}
	return *myorder
}
func (o order) getid() string {
	return o.id
}
func main (){
	order:=order{
		id:"1",
		amount:50.20,
		status:"ok",
		cretatedat:time.Now(),



	}
	my:=neworder("5",50.2,"completed")
	fmt.Printf("Order: ", order)
	order.changestatus("delivered")
	fmt.Println("Order: ", order)
	fmt.Println(my.getid())

}