package main

import "fmt"

//any and interface{} are same and can be used interchangeably and not suggested to use interface{} as it is old way of doing things


func printSlice[T int | string](s []T){
	for _,i:=range s{
		fmt.Println(i)
	}
}

// func printSliceString(s []string){
// 	for _,i:=range s{
// 		fmt.Println(i)
// 	}
// }


func main(){
	printSlice([]int{1,2,3,4,5})
	names:=[]string{"satyam","sharma","hello"}
	printSlice(names)

}