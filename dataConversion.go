package main

import "fmt"

func main() {
	var value32 int32 = 100000
	var value64 int64 = int64(value32)
	var value8 int8 = int8(value32)

	fmt.Println(value32)
	fmt.Println(value64)
	fmt.Println(value8) //maximum variable size exceed

	var name = "Geraldie"
	var bt byte = name[0]
	var btString string = string(bt)

	fmt.Println(name)
	fmt.Println(bt)
	fmt.Println(btString)
}
