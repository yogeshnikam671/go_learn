package main

import "fmt"

func main() {
	fmt.Println("type stringer error\n")
	typeAssertions()
}

// Type of a variable `a` can be asserted using the syntax : a.(<type>)
// This expression returns two things. The `a` value itself and an `ok` flag.
// If we consume the `ok` flag and `a` did not match the type, the `ok` flag value will be false and there won't be any error
// Whereas if we do not consume `ok`, it will error out.
// This works only for interfaces. To make it work for primitives, we have to declare primitives as blank interfaces
func typeAssertions() {
	fmt.Println("Type assertions only work on interfaces in go")
	fmt.Println("If we want it to work for primitives, declare primitive variables as type blank interface")
	var a interface{} = "this is a string"
	typeCheckedA := a.(string)
	fmt.Println("\nType checking the correct type without `ok` does not throw any error: " + typeCheckedA)

	_, ok := a.(int)
	fmt.Print("Invalid Type Checked `ok` value: ")
	fmt.Println(ok)

	// This will throw error, I have already explained why
	// invalidTypeCheckedA := a.(float64)
	// fmt.Println(invalidTypeCheckedA)
}
