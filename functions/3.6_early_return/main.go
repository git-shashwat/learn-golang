package main

import (
	"errors"
	"fmt"
)

func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("cannot divide by 0")
	}
	return dividend / divisor, nil
}

func main() {
	_, err := divide(5, 0)
	if err != nil {
		fmt.Println(err)
	}

	quotient, _ := divide(15, 3)
	fmt.Println("result", quotient)
}
