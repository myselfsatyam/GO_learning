package main

import "fmt"

//slices are dynamic array , used when size is not known
//most useful construct in go

func main(){
   //declaring a slice
   var nums []int //unitialised slices are nil
   fmt.Println(nums == nil)
   fmt.Println(len(nums))

   //make is used to make slice not nil during intializing
   //capacity->maximum number of elemts it can fit
   var a=make([]int,2,5) //it takes mainly 3 parameter datatype(req),no. of elem(req),capacity(optional)
   fmt.Println(nil==a)
    //when number of elements became equal to capacity, then capacity is doubled
	a[0]=1
   a=append(a, 2)
   a=append(a, 3)
   a=append(a, 4)
   a=append(a, 5)
   fmt.Println(a,len(a),cap(a))
   a=append(a, 6)
   fmt.Println(a,len(a),cap(a))
}
