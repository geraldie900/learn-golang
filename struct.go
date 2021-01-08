package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

//Struct method
func (cust Customer) sayHi(name string) {
	fmt.Println("Hello", cust.Name, ", my name is", name)
}

func main() {
	var ger Customer
	ger.Name = "Geraldie"
	ger.Address = "Indonesia"
	ger.Age = 20
	fmt.Println(ger.Name)
	fmt.Println(ger.Address)
	fmt.Println(ger.Age)

	tan := Customer{
		Name:    "Tanu",
		Address: "Lampung",
		Age:     23,
	}
	fmt.Println(tan)

	sap := Customer{"Saputra", "Tangerang", 30}
	fmt.Println(sap)

	ger.sayHi("Aldie")
}
