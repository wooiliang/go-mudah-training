// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

import (
	"fmt"
)

// Add imports.

// Add user type and provide comment.
// name, email, address

type user struct {
	name    string
	email   string
	address string
}

func main() {

	// Declare variable of type user and init using a struct literal.
	user := user{
		name:    "John",
		email:   "john@gmail.com",
		address: "1 Montreal Drive",
	}

	// Display the field values.
	fmt.Println(user.name)
	fmt.Println(user.email)
	fmt.Println(user.address)

	// Declare a variable using an anonymous struct.
	user2 := struct {
		name    string
		email   string
		address string
	}{
		name:    "Peter",
		email:   "peter@gmail.com",
		address: "1 Happy Land",
	}

	// Display the field values.
	fmt.Println(user2.name)
	fmt.Println(user2.email)
	fmt.Println(user2.address)
}
