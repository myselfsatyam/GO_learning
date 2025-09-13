package main

import "fmt"

//pass by value(a distinct copy) so instead of this use pass by refrence
// func changeNum(num int){
// 	num=5
// 	fmt.Println("in change num",num)
// }

//pass by refrence
func changeNum(num *int){
	*num=5 //derefencing
	fmt.Println("in change num",num)
	fmt.Println("in change num",*num)
}


func main(){
	num:=1
	//changeNum(num)
	//fmt.Println("After changeNum in main", num)
	fmt.Println("Memory Address", &num)

	changeNum(&num)
	fmt.Println("After changeNum in main", num)


}