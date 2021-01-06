package main

import "fmt"

func main() {
	// const must be declared immediatelly
	const firstname = "Geraldie"
	const lastname = "Tanu Saputra"
	const value = 500

	//const doesn't give error if the variable doesn't used
	//const can't be changed
	//value = 1000

	fmt.Println(firstname)
	fmt.Println(lastname)
	fmt.Println(value)

	const (
		country  string = "Indonesia"
		currency        = "Rupiah"
		nominal         = 1000
	)

	fmt.Println(country)
	fmt.Println(currency)
	fmt.Println(nominal)

}
