package main

import "fmt"

func getExpenseReport(e expense) (source string, receiver string, cost float64) {
	switch v := e.(type) {
	case email:
		return "email", v.toAddress, v.cost()
	case sms:
		return "sms", v.toPhoneNumber, v.cost()
	default:
		return "invalid", "", 0.0
	}
}

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	toAddress    string
	body         string
}

func (e email) cost() float64 {
	if e.isSubscribed {
		return float64(len(e.body)) * .01
	}
	return float64(len(e.body)) * .05
}

type sms struct {
	isSubscribed  bool
	toPhoneNumber string
	body          string
}

func (s sms) cost() float64 {
	if s.isSubscribed {
		return float64(len(s.body)) * .01
	}
	return float64(len(s.body)) * .03
}

func test(e expense) {
	source, receiver, cost := getExpenseReport(e)
	fmt.Printf("The %s goint to %v is going to cost %.1f\n", source, receiver, cost)
}

func main() {
	test(email{
		isSubscribed: true,
		toAddress:    "abc@gmail.com",
		body:         "hello world",
	})
	test(email{
		isSubscribed: false,
		toAddress:    "abc@gmail.com",
		body:         "hello world",
	})
	test(sms{
		isSubscribed:  true,
		toPhoneNumber: "000000000",
		body:          "hello world",
	})
	test(sms{
		isSubscribed:  false,
		toPhoneNumber: "000000000",
		body:          "hello world",
	})
}
