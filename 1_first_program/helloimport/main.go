package main

import (
	"fmt"

	"github.com/erikreppel/go-tutorial/1_first_program/hellopackage"
)

func main() {
	// := indicates type inference, greetings is a string
	greeting := hellopackage.Greeting("SDAML")
	// ^ this is the same as this v
	// var greeting string = hellopackage.Greeting("SDAML")
	fmt.Println(greeting)
}
