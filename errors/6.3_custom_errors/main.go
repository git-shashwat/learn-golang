package main

import (
	"fmt"
)

type divideError struct {
	dividend float64
}

func (e divideError) Error() string {
	return fmt.Sprintf("cannot divide %f by zero", e.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}

func main() {
	quotient, _ := divide(20, 4)
	fmt.Println("result:", quotient)

	_, err := divide(5, 0)
	if err != nil {
		fmt.Println(err)
	}
}
