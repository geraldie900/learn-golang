package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Geraldie",
		"address": "Tangerang",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Your name"
	book["author"] = "Geraldie"
	book["year"] = "2020"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "year")

	fmt.Println(book)
	fmt.Println(len(book))
}
