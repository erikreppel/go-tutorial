package main

import (
	"fmt"
)

// ErrorStruct is a stuct that can be used as an error
type ErrorStruct struct {
	Name  string
	Value string
}

func (e *ErrorStruct) Error() string {
	return "error, this stuct has name and value of value " + e.Name + e.Value
}

// ErrorInt is an int that fufils the error interface
type ErrorInt int

func (e ErrorInt) Error() string {
	return fmt.Sprintf("int %d had an error", e)
}

// We can even use functions to fufil interfaces

// Define a type of function that takes and int and adds an int
type simpleErrorFunc func(int) int

// Have that implement the interface
func (sf simpleErrorFunc) Error() string {
	return "Yes, even functions can be errors"
}

// Add implements the function signature of simpleFunc
func Add(x int) int {
	return x + 1
}

// then cast Add as a simpleErrorFunc

// SimpleErrorAdd implements the inferface error
var SimpleErrorAdd = simpleErrorFunc(Add)

// ErrorCheck verifies something is an error and calls Error()
func ErrorCheck(err error) {
	fmt.Println(err)
}

func main() {
	es := &ErrorStruct{Name: "Erik", Value: "asgfasg"}
	ErrorCheck(es)

	ei := ErrorInt(5)
	ErrorCheck(ei)

	ErrorCheck(SimpleErrorAdd)

}
