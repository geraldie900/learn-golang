package main

import "fmt"

func main() {
	var name string

	name = "Geraldie"
	fmt.Println(name)

	name = "Tanu Saputra"
	fmt.Println(name)

	//write type of variable is optional, if the variable is initialize immediately
	var str = "this is string"
	fmt.Println(str)

	//default type for number is int32
	var age = 30
	fmt.Println(age)

	//write var of variable is optional, but it changes the way of write variable, and assigned to default type data
	language := "English"
	fmt.Println(language)
	//using := only for first declaration

	var (
		firstname = "Geraldie"
		lastname  = "Tanu Saputra"
	)

	fmt.Println(firstname)
	fmt.Println(lastname)
}
