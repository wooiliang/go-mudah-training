// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare and make a map of integer values with a string as the key. Populate the
// map with five values and iterate over the map to display the key/value pairs.
package main

import (
	"fmt"
)

// Add imports.

func main() {

	// Declare and make a map of integer type values.
	a := make(map[string]int)

	// Initialize some data into the map.
	a["a"] = 1
	a["b"] = 2
	a["c"] = 3
	a["d"] = 4
	a["e"] = 5

	// Display each key/value pair.
	for key, value := range a {
		fmt.Println(key, value)
	}
}
