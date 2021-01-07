package main

import "fmt"

type blcklist func(string) bool

func registerUser(name string, blacklist blcklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome,", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("Geraldie", blacklist)
	//or
	registerUser("root", func(name string) bool {
		return name == "root"
	})
}
