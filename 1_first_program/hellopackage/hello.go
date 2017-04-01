package hellopackage

// Greeting greets people.
// Public functions in go start with a capital letter, private with lower case.
// Public functions should have a description of the function in a comment above
// the declaration for easy godoc generation.
func Greeting(name string) string {
	return "Hello " + name
}
