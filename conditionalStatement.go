package main

import "fmt"

func main() {
	name := "Tanu"

	if name == "Geraldie" {
		fmt.Println("Hello,", name)
	} else if name == "Tanu" {
		fmt.Println("Hello,", name)
	} else {
		fmt.Println("What's your name?")
	}

	//Short statement
	lastname := "tanu saputra"

	if lth := len(lastname); lth > 3 {
		fmt.Println("Too long?")
	} else {
		fmt.Println("Too short?")
	}

	fmt.Println("====================================")

	name = "Geraldie"

	switch name {
	case "Geraldie":
		fmt.Println("Hello, Ger")
	case "Tanu":
		fmt.Println("hello, Tan")
	case "Saputra":
		fmt.Println("Hello, Putra")
	default:
		fmt.Println("Who the f*ck are you?")
	}

	switch lth := len(name); lth > 5 {
	case true:
		fmt.Println("Woaw")
	case false:
		fmt.Println("Roar")
	}

	lth := len(name)
	switch {
	case lth > 10:
		fmt.Println("Omg?")
	case lth < 10:
		fmt.Println("Wew?")
	}
}
