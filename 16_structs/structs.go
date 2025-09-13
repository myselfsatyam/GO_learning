package main

import (
	"fmt"
	"time"
)

//Go struct ↔ TS type / object literal

type customer struct{
	name string
	phone string
}

type order struct{ //it acts as blueprint
	id string
	amount float32
	status string
	createdAT time.Time //nanp second precision
	//struct embedding
	customer

}

func newOrder(id string,amount float32,status string) *order{ //constructor for creating new order
	myOrder:=order{
		id:id,
		amount: amount,
		status: status,
	}
	return &myOrder
}

func (o *order) changeStatus(status string){ //* is used at the time of mofifying a value
    o.status=status
}

func (o order) getAmount() float32{
	return o.amount
}

func main(){

   language := struct{
     name string
	 isgood bool
   }{"golang",true}

   fmt.Println(language)
   


	myorder:=order{
		id:"1",
		amount: 13.00,
		status: "recieved",
		createdAT: time.Now(),
	}

	myorder2:=order{
		id: "2",
		amount: 1200,
		status: "inprogress",
	}

	//myorder3:=newOrder("3",1234,"confirmed")
	myorder3:=*newOrder("3",1234,"confirmed")

	myorder.createdAT=time.Now()

	//fmt.Println(myorder.status)
	//myorder.status="on way"

	//if field is not set then deafult value is zero value for each data structure

	myorder.changeStatus("confirmed")

	fmt.Println("Order struct ",myorder)
	fmt.Println("Order struct ",myorder2)
	fmt.Println("order struct",myorder3)

	//struct embedding usage
newCustomer:=customer{
	name: "jone",
	phone: "1231ec",
}

	myorder4:=order{
		id:"1",
		amount: 1200,
		status: "done",
		customer: newCustomer,

	}

	fmt.Println(myorder4)
	myorder4.customer.name="robin"
	fmt.Println(myorder4)
	fmt.Println(myorder4.getAmount())
	


}