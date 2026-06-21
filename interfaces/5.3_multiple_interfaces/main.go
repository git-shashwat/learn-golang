package main

import "fmt"

type expense interface {
	cost() float64
}

type printer interface {
	print()
}

type email struct {
	isSubscribed bool
	body         string
}

func (e email) cost() float64 {
	if e.isSubscribed {
		return float64(len(e.body)) * .01
	}
	return float64(len(e.body)) * .05
}

func (e email) print() {
	fmt.Println(e.body)
}

func test(e expense, p printer) {
	p.print()
	fmt.Printf("cost of email is %.1f\n", e.cost())
}

func main() {
	e := email{
		isSubscribed: true,
		body:         "hello world",
	}
	test(e, e)

	e = email{
		isSubscribed: false,
		body:         "hello world",
	}
	test(e, e)
}
