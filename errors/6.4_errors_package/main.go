package main

import (
	"errors"
	"fmt"
)

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("no dividing by 0")
	}
	return x / y, nil
}

func main() {
	quotient, _ := divide(20, 4)
	fmt.Printf("quotient: %f", quotient)

	fmt.Println()

	_, err := divide(5, 0)
	fmt.Println(err)
}
