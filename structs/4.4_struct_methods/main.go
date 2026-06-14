package main

import "fmt"

type authInfo struct {
	username string
	password string
}

func (a authInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", a.username, a.password)
}

func main() {
	userAuth := authInfo{
		username: "shashwat",
		password: "123456",
	}

	fmt.Println(userAuth.getBasicAuth())
}
