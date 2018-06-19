// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a nil slice of integers. Create a loop that appends 10 values to the
// slice. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main

// Add imports.
import "fmt"

func main() {

	// Declare a nil slice of integers.
	var a []int

	// Appends numbers to the slice.
	for index := 0; index < 10; index++ {
		a = append(a, index*10)
	}

	// Display each value in the slice.
	for _, value := range a {
		fmt.Println(value)
	}

	fmt.Println("")

	// Declare a slice of strings and populate the slice with names.
	// b := make([]string, 5, 5)
	// b[0] = "a"
	// b[1] = "b"
	// b[2] = "c"
	// b[3] = "d"
	// b[4] = "e"
	b := []string{"a", "b", "c", "d", "e"}

	// Display each index position and slice value.
	for i, value := range b {
		fmt.Println(i, value)
	}

	fmt.Println("")

	// Take a slice of index 1 and 2 of the slice of strings.
	c := b[1:3]

	// Display each index position and slice values for the new slice.
	for i, value := range c {
		fmt.Println(i, value)
	}
}
