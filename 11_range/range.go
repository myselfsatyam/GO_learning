package main

import "fmt"

// range is mainly used with for loop, for iterating over data structure

func main(){
	 nums:=[]int {6,7,8}
	 sum:=0
    //iterating with traditional for loop
	//  for i:=0;i<len(nums);i++{
	// 	fmt.Println(nums[i])
	//  }

	 for _,num:=range nums{
		//fmt.Println(num)
		sum=sum+num
	 }
	 fmt.Println(sum)

	 for i,num:=range nums{
		fmt.Println(i,num)
	 }

	 m:= map[string]string{"fname":"go","lname":"lang"}
	 for k,v:=range m{
		fmt.Println(k,v)
	 }

	 for _,v:=range m{
		fmt.Println(v)
	 }
	 for k:=range m{
		fmt.Println(k)
	 }

	 //range on string

	 for i,c:=range "gp"{
		fmt.Println(i,c) //i gives starting byte of rune and c gives unicode(uncicode code point to rune)
	 }

}
