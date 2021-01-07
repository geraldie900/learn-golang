package main

import "fmt"

func sayHello() {
	fmt.Println("Hello World")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello,", firstName, lastName)
}

func sayHelloReturn(name string) string {
	return "Hey," + name
}

func sayHelloMultiReturn() (string, string) {
	return "Geraldie", "Tanu Saputra"
}

//named return function
func getFullName() (first, middle, last string) { //or string first, string middle, string last
	first = "ger"
	middle = "tan"
	last = "sap"
	return //no need to write variable name, because already declared in first with func declaration
}

//variadic function
//variadic argument only and must declared in last position of paramater
//pamater may be empty
func sumAll(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}
	return total
}

func goodBye(name string) string {
	return "Good Bye " + name + "!"
}

func main() {
	sayHello()
	sayHelloTo("Geraldie", "Tanu Saputra")
	result := sayHelloReturn("Geraldie")
	fmt.Println(result)
	first, last := sayHelloMultiReturn()
	fmt.Println("this is multi function :", first, last)
	//ignore second return value
	first1, _ := sayHelloMultiReturn()
	fmt.Println("this is multi function, but ignore second value :", first1)

	a, b, c := getFullName()
	fmt.Println(a, b, c)
	fmt.Println(sumAll(10, 20, 30, 40, 50))

	slice := []int{10, 20, 30, 40, 50}
	total := sumAll(slice...)
	fmt.Println(total)

	sayGoodBye := goodBye //kind like giving alias to a function
	result = sayGoodBye("geraldie")
	fmt.Println(result)

}
