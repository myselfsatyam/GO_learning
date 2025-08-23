//maps is associatve data sturcture
//map[key]value

package main

import (
	"fmt"
	"maps"
)

func main(){
	//creating a map
	m2:=map[string]int{"price":40,"age":10}
	fmt.Println(m2)
	//creating a map using make
	m:=make(map[string]string)
	//setting a elemen
	m["name"]="Golang"
	m["usecase"]="gprc"
	fmt.Println(m)
	fmt.Println(m["name"])
	//accesing a undeclared key (if key value pair doesnot exist it return falsy value)
	fmt.Println(m["phone"]) //m is of string type it return empty

	m1:=make(map[string]int)
	m1["age"]=30
	fmt.Println(m1["age"])
	fmt.Println(m1["phone"])
	//lenght of map
	fmt.Println(len(m),len(m1))

	//deleting a key value pair from map
	delete(m,"usecase")
	fmt.Println(m)

	//clearing entire map
	clear(m)
	fmt.Println(m)

	//checking for particular elemt present in map
	v,ok:=m2["price"]
    fmt.Println(v)
	if ok {
		fmt.Println("All ok")
	} else{
		fmt.Println("not ok")
	}

	//comparing 2 maps using maps package

	fmt.Println(maps.Equal(m1,m2))


}