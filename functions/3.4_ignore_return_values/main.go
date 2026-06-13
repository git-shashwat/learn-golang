package main

import "fmt"

func getNames() (string, string) {
	return "shashwat", "tyagi"
}

func main() {
	firstName, _ := getNames()
	fmt.Println("hello", firstName)
}
