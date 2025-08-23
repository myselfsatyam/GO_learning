package main

import (
	"fmt"
	"slices"
)

//slices are dynamic array , used when size is not known
//most useful construct in go

func main(){
   //declaring a slice
   var nums []int //unitialised slices are nil
   fmt.Println(nums == nil)
   fmt.Println(len(nums))

   //make is used to make slice not nil during intializing
   //capacity->maximum number of elemts it can fit
   var a=make([]int,0,5) //it takes mainly 3 parameter datatype(req),no. of elem(req),capacity(optional)
   fmt.Println(nil==a)
    //when number of elements became equal to capacity, then capacity is doubled
   a=append(a, 1)
   a=append(a, 2)
   a=append(a, 3)
   a=append(a, 4)
   a=append(a, 5)
   fmt.Println(a,len(a),cap(a))
   a=append(a, 6)
   fmt.Println(a,len(a),cap(a))

   //shorthand for slice
   b:=[]int{}
   b=append(b, 1)
   b=append(b, 2)
   fmt.Println(b,len(b),cap(b))

   //slice using index
   var c=make([]int,2,5)
   c[0]=2
   c[1]=4
   fmt.Println(c,len(c),cap(c))

   //slices are copied using copy function
   var n1=make([]int,0,2)
   n1=append(n1,2)
   var n2=make([]int,len(n1))

   fmt.Println(n1,n2)
   //copying n1 to n2
   copy(n2,n1) //copy(dest,source)
    fmt.Println(n1,n2)


	//slice operator(slicing)
	var n3=[]int{1,2,3}
	fmt.Println(n3[0:2])
	fmt.Println(n3[:])
	fmt.Println(n3[0:])
	fmt.Println(n3[:2])

	//slices package
	var a1=[]int{1,2,4}
	var a2=[]int{1,2,3}

    fmt.Println(slices.Equal(a1,a2))

	//2-D slices
	var d=[][]int{{1,2},{3,5}}
	fmt.Println(d)
}
