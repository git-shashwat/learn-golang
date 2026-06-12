package main

import "fmt"

func main() {
	var smsLimit int = 100
	var costPerMessage float32 = .02
	var hasPermissions bool = true
	var username string = "shashwat"

	fmt.Printf("%v %.2f %v %q\n", smsLimit, costPerMessage, hasPermissions, username)
}
