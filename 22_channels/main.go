package main

import (
	"fmt"
	"time"
)

//channels are used to communicate between goroutines
//they provide a way to send and receive values between goroutines

func processNum(numChan chan int){
	fmt.Println("Processing number:", <-numChan)

}
//recieving and sending values through channels is blocking
//meaning that the goroutine will wait until the value is sent or received
func sum(result chan int,num1 int,num2 int){
     result<- num1+num2 //sending the sum of num1 and num2 to the result channel
}


//go routine synchronizer
//unbuffered channels(one value at a time) can be used to synchronize goroutines
//the main goroutine can wait for another goroutine to finish by receiving a value from an unbuffered channel
func taskdone(done chan bool){
	defer func() {done<-true}() //send a value to the done channel when the function exits
	fmt.Println("Task is processing...")

}


//buffered channels
//channels can be buffered, meaning they can hold a certain number of values without a corresponding receiver for those values
//this allows goroutines to send values into the channel without blocking until the buffer is full
//make(chan type, capacity)
//if the buffer is full, the sending goroutine will block until a value is received from the channel
//if the buffer is empty, the receiving goroutine will block until a value is sent into the channel
func emailSender(email chan string,done chan bool){
	defer func() {done<-true}() //signal that the email sending is done
	for emailAddr:= range email{
		fmt.Println("Sending email to:",emailAddr)
		time.Sleep(time.Second) //simulate time taken to send an email
	}
}



func main() {
	//messages:=make((chan string))

	//messages<-"ping" //send a value into the channel

	//msg:= <-messages ///receive a value from the channel
	//fmt.Println(msg) // this will cause a deadlock( deadlock means that the program is stuck and cannot proceed because it is waiting for something that will never happen. In this case, the main goroutine is trying to receive a value from the channel, but there are no other goroutines that are sending values into the channel. As a result, the main goroutine is blocked and cannot continue executing.)
 

	numChan:=make(chan int)
	go processNum(numChan)
	numChan<-42

	time.Sleep(time.Second * 2)

	result:=make(chan int)
	go sum(result,5,7)
	fmt.Println("Sum:",<-result)

	done:=make(chan bool)
	go taskdone(done)
	<-done //wait for the task to be done(blocking)
	fmt.Println("Task is done")

	 email:=make(chan string,100)//buffered channel with a capacity of 100
	 go emailSender(email,done)
	//sending multiple emails to the channel
	// email<-"sam@gmail.com"
	// email<-"sa@gmail.com"

	for i:=0;i<10;i++{
         email<- fmt.Sprintf("%d@gmail.com",i)
	}
	fmt.Println("All emails sent to the channel")
    close(email) //close the email channel to signal that no more emails will be sent
	<-done //wait for the email sender to finish
	fmt.Println("Email sender is done")


}
	