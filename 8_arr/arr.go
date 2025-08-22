package main

import (
	"fmt"
)

//array are numbeered sequnece of specific length

//array are used when there is fixed size 
//memory optimization
//fixed time
func main(){
//declaring an array
  var nums [5]int
  nums[0]=1 //intialising value at 0th index , o is falsy value for int
  fmt.Println(nums)

  //length of array
    //fmt.Println(len(nums))
  
  var bool [4]bool
  bool[0]=true
  fmt.Println(bool) //false is falsy value for boolean

  var name [3]string
  name[0]="golang"
  fmt.Println(name) //empty spaces are falsy va;ues for string

  //int->0,bool->false,string=""

  //intialising values while declaring
  num:=[3]int{1,2,3}
  fmt.Println(num)


  //2-d array in go
  nd:=[2][2]int{{1,2},{3,4}}
   fmt.Println(nd)
}
