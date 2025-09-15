package main

import (
	"fmt"
	"os"
)


func main(){

	//open a file
	// f,err:=os.Open("example.txt")
	// if err !=nil{
	// 	//log the error and exit
	// 	panic(err)
	// }

	// fileinfo,err:=f.Stat()
	// if err!=nil{
	// 	panic(err)
	// }

	// fmt.Println("File name:",fileinfo.Name())
	// fmt.Println("file or folder:",fileinfo.IsDir())
	// fmt.Println("File size in bytes:",fileinfo.Size())
	// fmt.Println("File mode:",fileinfo.Mode())
	// fmt.Println("File modification time:",fileinfo.ModTime())



	//read the file
	f,err:=os.ReadFile("example.txt")
	if err!=nil{
		panic(err)
	}
	fmt.Println(string(f)) //convert the byte slice to a string and print it

	//write to a file
	err=os.WriteFile("example2.txt",f,0644) //0644 is the file permission
	if err!=nil{
		panic(err)
	}

	//get file info
	fileinfo,err:=os.Stat("example2.txt")
	if err!=nil{
		panic(err)
	}

	fmt.Println("File name:",fileinfo.Name())
	fmt.Println("file or folder:",fileinfo.IsDir())
	fmt.Println("File size in bytes:",fileinfo.Size())
	fmt.Println("File mode:",fileinfo.Mode())
	fmt.Println("File modification time:",fileinfo.ModTime())

	//delete a file
	err=os.Remove("example2.txt")
	if err!=nil{
		panic(err)
	}
	fmt.Println("File deleted successfully")

	//rename a file
	err=os.Rename("example.txt","example_renamed.txt")
	if err!=nil{
		panic(err)
	}
	fmt.Println("File renamed successfully")



	
}