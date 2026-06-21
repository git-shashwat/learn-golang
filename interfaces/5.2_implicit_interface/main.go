package main

import "fmt"

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

type fullTime struct {
	name   string
	salary int
}

func (fT fullTime) getName() string {
	return fT.name
}

func (fT fullTime) getSalary() int {
	return fT.salary
}

func test(emp employee) {
	fmt.Printf("Pay for %s is %v\n", emp.getName(), emp.getSalary())
}

func main() {
	test(contractor{
		name:         "shashwat",
		hourlyPay:    20,
		hoursPerYear: 100,
	})
	test(fullTime{
		name:   "tyagi",
		salary: 60000,
	})
}
