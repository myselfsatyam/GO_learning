package main

import "fmt"

//In Go, variadic functions are functions that can take a variable number of arguments. They’re useful when you don’t know beforehand how many inputs you’ll need to handle.


func sum(nums ...int){
	fmt.Println("Numbers:",nums)
	total:=0

	for _,v:= range nums{
		total=total+v
	}
	fmt.Println("Sum:",total)

}


func main(){
	sum(1,2)
	sum(12,3,4,5,7)

}