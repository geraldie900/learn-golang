package main

import "fmt"

func main() {
	var test = 90
	var attendance = 10

	var passTest = test > 75
	var passAttendance = attendance > 5

	var pass = passTest && passAttendance
	fmt.Println(pass)
}
