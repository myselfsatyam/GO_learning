package main

import "fmt"

//const using grouping
const (
	port=5000
	host="localhost"
)

const owner="Nightwing aka Robin"
func main(){
	//constants are immutable i.e. cannot be change, constants are declared using const keyword
 const name="golang" //constant are read during complile time
 const age=30

 //const are mainly described used outside the function

 fmt.Println(port,host)

}