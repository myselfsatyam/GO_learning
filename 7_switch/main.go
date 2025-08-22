package main

import (
	"fmt"
	"time"
)

func main() {
	//simple switch

	i := 5
// don't need to right break after each case

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other") //default is optional		
	}	

	
	//multiple condition switch

	switch time.Now().Weekday(){
	case time.Saturday,time.Sunday:
		fmt.Println("It's Weekend")
	default:
		fmt.Println("It's Weekday")

		 
	}

	//type switch
	whoAmI:=func(i interface{}){
		switch i.(type){
		case int:
			fmt.Println("it's a integer")
		case string:
			fmt.Println("string")
		case bool:
			fmt.Println("boolean")	
		default:
			fmt.Println("other")		
		}
	}

	whoAmI("true")
}
