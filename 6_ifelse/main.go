package main

import "fmt"

func main(){
	age:=16

	if age>18{
		fmt.Println("You are elgibile to vote")
	} else if age==18{
         fmt.Println("u are exception")
	} else{
		fmt.Println("inelgibile to vote")
	}
//go doesnot have ternary operator we use of else
}