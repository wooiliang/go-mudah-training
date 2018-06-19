// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).
package main

// Add imports
import "fmt"

// main is the entry point for the application.
func main() {

	// Declare variables that are set to their zero value.
	var a int
	var b string
	var c float64
	var d bool

	// Display the value of those variables.
	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var d bool \t %T [%v]\n\n", d, d)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	aa := 1
	bb := "Hello World"
	cc := 0.1
	dd := true

	// Display the value of those variables.
	fmt.Printf("aa := 1 \t %T [%v]\n", aa, aa)
	fmt.Printf("bb := \"Hello World\" \t %T [%v]\n", bb, bb)
	fmt.Printf("cc := 0.1 \t %T [%v]\n", cc, cc)
	fmt.Printf("dd := true \t %T [%v]\n\n", dd, dd)

	// Perform a type conversion.
	aaa := int64(99)

	// Display the value of that variable.
	fmt.Printf("aaa := int64(99) \t %T [%v]\n", aaa, aaa)
}
