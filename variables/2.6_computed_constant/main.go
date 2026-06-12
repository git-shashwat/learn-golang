package main

import "fmt"

func main() {
	const secondsInMinute = 60
	const minutesInHour = 60

	const secondsInHour = minutesInHour * secondsInMinute

	fmt.Println("Seconds in hour:", secondsInHour)
}
