package main

import "fmt"

type messageToSend struct {
	message  string
	sender   user
	receiver user
}

type user struct {
	name        string
	phoneNumber int
}

func canSendMessage(m messageToSend) bool {
	if m.sender.name != "" && m.sender.phoneNumber != 0 && m.receiver.name != "" && m.receiver.phoneNumber != 0 {
		return true
	}

	return false
}

func main() {
	s1 := user{
		name:        "st",
		phoneNumber: 9807654321,
	}
	r1 := user{
		name:        "bs",
		phoneNumber: 9807653421,
	}
	sms := messageToSend{
		message:  "ily",
		sender:   s1,
		receiver: r1,
	}

	canSendSMS := canSendMessage(sms)
	fmt.Println("can s1 send message to r1", canSendSMS)

	s1 = user{
		name: "st",
	}
	r1 = user{
		phoneNumber: 9807653421,
	}
	sms = messageToSend{
		message:  "ily",
		sender:   s1,
		receiver: r1,
	}

	canSendSMS = canSendMessage(sms)
	fmt.Println("can s1 send message to r1", canSendSMS)
}
