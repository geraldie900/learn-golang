package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Geraldie"
	names[1] = "Tanu"
	names[2] = "Saputra"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		100,
		95,
		90,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	//len only return capacity of the array, not return how many data already in array
	fmt.Println(len(names))
	fmt.Println(len(values))

	var arr [10]int
	fmt.Println(len(arr))
}
