package main

import "fmt"

func main() {
	const messageLen = 10
	const maxMessageLen = 20

	if messageLen <= maxMessageLen {
		fmt.Println("message sent")
	} else {
		fmt.Println("message not sent")
	}
}
