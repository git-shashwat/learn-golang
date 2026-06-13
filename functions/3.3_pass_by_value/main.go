package main

import "fmt"

func incrementSent(s int, amt int) int {
	s += amt
	return s
}

func main() {
	smsSentSoFar := 420
	smsToSend := 25

	fmt.Println("sms sent so far:", smsSentSoFar)

	fmt.Println("sending another", smsToSend, "messages")
	smsSentSoFar = incrementSent(smsSentSoFar, smsToSend)

	fmt.Println("sms sent so far:", smsSentSoFar)
}
