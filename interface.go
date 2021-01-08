package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func sayHello(hasName HasName) {
	fmt.Println("Hey,", hasName.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

//Empty Interface
func oops(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "oops"
	}
}

func main() {
	var ger Person
	ger.Name = "Geraldie"
	sayHello(ger)

	cat := Animal{
		Name: "cat",
	}
	sayHello(cat)

	fmt.Println("==================")
	var data interface{} = oops(1)
	fmt.Println(data)
}
