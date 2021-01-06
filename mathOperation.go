package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 10
	var c = a + b
	fmt.Println(c)

	var x = 10
	x += 10
	fmt.Println(x)

	x++
	fmt.Println(x)

	var negative = -100
	var positive = +100
	fmt.Println(negative)
	fmt.Println(positive)
}
