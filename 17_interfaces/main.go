package main

import "fmt"


type paymenter interface {
     Payment(amount float32)
}

type payment struct {
    // gateway stripe\razorpay
	//gateway razorpay

    gateway paymenter
}

// we are violating open close pricniple here
// because if we want to add another payment gateway we need to modify this method
func (p payment) makePayment(amount float32){
   //razorpayPaymentGw:= razorpay{}
   //razorpayPaymentGw.Payment(amount)
   //stripePaymentGw:= stripe{}
   //p.gateway.pay(amount)
   p.gateway.Payment(amount)
}

type razorpay struct {

}

func (r razorpay) Payment(amount float32){
   
	fmt.Println("Payment done using razorpay")
}

type stripe struct {

}

func (s stripe) pay(amount float32){
   
	fmt.Println("Payment done using stripe")
}

type fakepayment struct {}

func (f fakepayment) Payment(amount float32){
    fmt.Println("Payment done using fakepayment for testing purpose")
}

type paypal struct {}   

func (p paypal) Payment(amount float32){
    fmt.Println("Payment done using paypal")
}

func main(){
    //stripePaymentGw:= stripe{}
	razorpayPaymentGw:= razorpay{}
	userpayment:= payment{gateway: razorpayPaymentGw}
	userpayment.makePayment(1000.00)

     fakeGw:= fakepayment{}
    newpayment:= payment{gateway: fakeGw}
     newpayment.makePayment(1000.00)

        paypalGw:= paypal{}
    paypalPayment:= payment{gateway: paypalGw}
     paypalPayment.makePayment(1000.00)
     
  
}
