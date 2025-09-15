package main

import (
	"fmt"

	"github.com/myselfsatyam/go_learning/25_packages/auth"
)

func main() {
	auth.Login("satyam","mypassword")
	session:=auth.GetsessionID()
	fmt.Print("Session ID:",session)
}