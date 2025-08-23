package main

import "fmt"



func add(a int,b int) int{ // for same parameter we can write as add(a,b int)
    return a+b
}

// mutliple return values return type are written in round braces
func getLang() (string,string,string){
	return "golang","js","ts"
}

//function accepting a function as parameter

	func processIt(fn func(a int) int)  {
		fn(1)
	}


func main(){
    x:=10
	y:=12
	result:=add(x,y)

	fmt.Println(result)
	lang1,lang2,lang3:=getLang()
	fmt.Println(lang1,lang2,lang3)

	lan1, _ , _ :=getLang()
	fmt.Println(lan1)

	fn:=func(a int) int{
		return 2
	}

	processIt(fn)


}