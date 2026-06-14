package main

import "fmt"

type messageToSend struct {
	phoneNumber int
	message     string
}

func test(m messageToSend) {
	fmt.Println("sending message:", m.message, "to", m.phoneNumber)
}

func main() {
	newMessage := messageToSend{
		phoneNumber: 1234567890,
		message:     "gulab jamun > rasmalai",
	}
	test(newMessage)

	newMessage = messageToSend{
		phoneNumber: 9087654321,
		message:     "hip hop > pop",
	}
	test(newMessage)
}
