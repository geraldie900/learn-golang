package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Loop -", counter)
		counter++
	}

	fmt.Println("===================================")

	for count := 1; count <= 10; count++ {
		fmt.Println("Loop -", count)
	}

	names := []string{"Geraldie", "Tanu", "Saputra"}
	for idx, name := range names {
		fmt.Println("Name on index", idx, "is", name)
	}

	//unused index, change to _
	for _, name := range names {
		fmt.Println("Name :", name)
	}

	person := make(map[string]string)
	person["name"] = "Geraldie"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, value)
	}
}
