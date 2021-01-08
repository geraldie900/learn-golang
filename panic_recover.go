package main

import "fmt"

func endApp() {
	message := recover()
	fmt.Println("Error with messsage :", message)

	fmt.Println("Application Finish Run")

}

//when panic triggered, the program will stop
//when there is recover, even panic triggered, the program still continue

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APPLICATION ERROR")
	}
	fmt.Println("Running App")
}

func main() {
	runApp(false)
}
