package main

import "fmt"

func main() {
	type NoKTP string //alias for string
	type Married bool

	var noKtp NoKTP = "123456789"
	var marriedStatus Married = true
	fmt.Println(noKtp)
	fmt.Println(marriedStatus)
}
