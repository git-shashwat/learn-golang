package main

import "fmt"

func concat(s1 string, s2 string) string {
	return s1 + s2
}

func main() {
	const name, greeting = "shashwat ", "hello"

	response := concat(name, greeting)

	fmt.Println(response)
}
