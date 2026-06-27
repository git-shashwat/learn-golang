package main

import (
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	mCCost, mCErr := sendSMS(msgToCustomer)
	if mCErr != nil {
		return 0, mCErr
	}

	mSCost, mSErr := sendSMS(msgToSpouse)
	if mSErr != nil {
		return 0, mSErr
	}

	return mCCost + mSCost, nil
}

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}

func main() {
	cost, err := sendSMSToCouple("hello world", "hello world")
	if err != nil {
		fmt.Println("failed to send message", err)
		return
	}
	fmt.Println("cost of sending text:", cost)
}
