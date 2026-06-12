package main

import "fmt"

func main() {
	var username string = "shashwat"
	// var password int = 2342342343
	var password string = "2342342343"

	// invalid operation of additon on string & int
	// fmt.Println("Basic Auth: ", username+":"+password)

	// to fix, change the type of password from int to string with the same value
	fmt.Println("Basic Auth: ", username+":"+password)
}
