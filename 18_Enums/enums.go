package main

import "fmt"

//enums -> enumarted types(go doesn't have enums but we can achieve similar functionality using iota)

type orderStatus int 

const (
	Received orderStatus = iota
	Confirmed
	Shipped
	Delivered
)

func changeOrderStatus(status orderStatus) {
	fmt.Println("Order status changed to ", status)
}

func main(){
    changeOrderStatus(Confirmed)
}
