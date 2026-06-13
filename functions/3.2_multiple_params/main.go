package main

import "fmt"

func createUser(firstName, lastName string, age int) {
	fmt.Println("user created:", firstName, lastName, "age:", age)
}

func main() {
	createUser("shashwat", "tyagi", 25)
}
