package main

import "fmt"

func factorial(value int) int {
	if value == 0 || value == 1 {
		return 1
	} else {
		return value * factorial(value-1)
	}
}

func main() {
	fact := factorial(5)
	fmt.Println(fact)
}
