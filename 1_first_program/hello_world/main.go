package main

import (
	"fmt"
)

type testStuct struct {
	Value int
}

func main() {
	fmt.Println("Hello, world!")

	fmt.Println("Go has structs and maps and slices(and arrays)")
	x := map[string]string{"hello": "world"}
	fmt.Println(x)

	z := []int{1, 2, 3}
	fmt.Println(z)
	fmt.Println("You can append to slices")
	z = append(z, 4)
	fmt.Println(z)

	fmt.Println("Struct:")
	y := testStuct{Value: 42}
	fmt.Println(y)
	fmt.Printf("There are multiple ways to print structs %+v\n", y)

	fmt.Println("Go also has pointers, but its cool")
	fmt.Printf("Pointer value for test struct %p\n", &y)

	fmt.Println("But theres no point arithmatic which is awesome!")

	fmt.Println("& gives you the address of something, * gives you the value at that address")

	yPointer := &y
	fmt.Println(y, *yPointer, "These are the same")

	fmt.Printf("Checkout %s for more examples of data types", "https://gobyexample.com")

}
