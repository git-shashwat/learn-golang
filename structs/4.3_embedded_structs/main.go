package main

import "fmt"

type user struct {
	name        string
	phoneNumber int
}

type sender struct {
	user
	rateLimit int
}

func test(s sender) {
	fmt.Println("sender info:", s.name, s.phoneNumber, s.rateLimit)
}

func main() {
	test(sender{
		user: user{
			name:        "shashwat",
			phoneNumber: 9876543210,
		},
		rateLimit: 500,
	})
}
