package main

//goroutines are lightweight threads managed by the Go runtime
//they allow concurrent execution of functions or methods without blocking the main program flow 

import (
	"fmt"
	"time"
)


func task(id int){
	fmt.Println("Task",id,"is being processed")
}

func main(){
	for i:=0;i<10;i++{
		//go task(i)

	go	func (i int)  {
			fmt.Println(i)
		}(i)
	}

	time.Sleep(time.Second * 2)

}