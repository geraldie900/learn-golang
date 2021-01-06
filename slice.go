package main

import "fmt"

func main() {
	var months = [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"November",
		"October",
		"December",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	//slice using preferance from array
	//months[5] = "Changed"
	//fmt.Println(slice1)

	slice1[0] = "MayAgain"
	fmt.Println(months)

	var days = [...]string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}
	var slice2 = days[5:]
	fmt.Println(slice2)

	//append will make new array, if capacity of array already exceeded
	var slice3 = append(slice2, "New-Day")
	fmt.Println(slice3)
	slice3[1] = "Not-Sunday"
	fmt.Println(slice3)
	//continue from last comment; this is why when the changing slice3, it doesn't changing the main array (days and slice2)
	fmt.Println(slice2)
	fmt.Println(days)
	fmt.Println("========================")

	newSlice := make([]string, 2, 5) // (type, len, cap)
	newSlice[0] = "Geraldie"
	newSlice[1] = "Tanu"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//make sure len of new slice is same as len copied slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	thisArray := [5]int{1, 2, 3, 4, 5} // or [...]
	thisSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(thisArray)
	fmt.Println(thisSlice)
}
