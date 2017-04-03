package main

import "fmt"
import "time"

func sayHi(name string) {
	fmt.Println("Hi " + name)
}

func main() {
	// Putting go infront of a function call runs it in a new go routine
	go sayHi("Jon")
	sayHi("Arya")
	go sayHi("Ned")

	// We need this otherwise the program will finish before goroutines can print
	time.Sleep(time.Second)
}
