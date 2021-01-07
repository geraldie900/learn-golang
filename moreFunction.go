package main

import "fmt"

type fltr func(string) string //make an alias for func func(string) string with name : Filter

//function as parameter
func sayHelloWithFilter(name string, filter fltr) { //or (name string, filter func(string) string)
	nameFiltered := filter(name)
	fmt.Println("Hello,", nameFiltered)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	}
	return name

}

func main() {
	sayHelloWithFilter("Geraldie", spamFilter)
	//or using alias
	filter := spamFilter
	sayHelloWithFilter("anjing", filter)
}
