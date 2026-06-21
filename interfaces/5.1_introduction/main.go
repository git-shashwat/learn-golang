package main

import (
	"fmt"
	"time"
)

func sendMessage(msg message) {
	fmt.Println(msg.getMessage())
}

type message interface {
	getMessage() string
}

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime)
}

type sendingReport struct {
	reportName      string
	numberOfReports int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your %s" report is ready. We have sent %v reports`, sr.reportName, sr.numberOfReports)
}

func test(m message) {
	sendMessage(m)
}

func main() {
	test(sendingReport{
		reportName:      "blood test",
		numberOfReports: 2,
	})
	test(birthdayMessage{
		birthdayTime:  time.Now(),
		recipientName: "shashwat",
	})
}
