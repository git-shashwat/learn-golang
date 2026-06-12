package main

import "fmt"

func main() {
	const name = "saul goodman"
	const openRate = 30.5

	msg := fmt.Sprintf("Hi %v, your open rate is %.1f percent", name, openRate)

	fmt.Println(msg)
}
