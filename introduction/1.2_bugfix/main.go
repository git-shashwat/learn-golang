package main

import "fmt"

func main() {
	const costPerMessage = .02
	const numberOfMessages = 200

	const totalCost = costPerMessage * numberOfMessages

	fmt.Printf("I spent $%.2f on messages today", totalCost)
}
