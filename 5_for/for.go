package main

import "fmt"

//for is the only loop in the golang

func main(){
//while loop using for
i:=1
for i<5{
	fmt.Println(i)
	i++
}

//infinte loop
// for {
// 	println("1")
// }


//traditionalfor loop

for i:=0;i<=3;i++{
	fmt.Println(i)
}

//for loop having range
for i:= range 11{
	fmt.Println(i) //prints from 0 to 10
}


}
