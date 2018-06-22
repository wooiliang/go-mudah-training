// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program that uses a fan out pattern to generate 100 random numbers
// concurrently. Have each goroutine generate a single random number and return
// that number to the main goroutine over a buffered channel. Set the size of
// the buffer channel so no send every blocks. Don't allocate more buffers than
// you need. Have the main goroutine display each random number is receives and
// then terminate the program.
package main

// Add imports.
import (
	"fmt"
	"math/rand"
	"time"
)

// Declare constant for number of goroutines .

func init() {
	// Seed the random number generator.
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// Create the buffer channel with a buffer for
	// each goroutine to be created.
	emps := 20
	ch := make(chan int, emps)

	// Iterate and launch each goroutine.
	for index := 0; index < emps; index++ {
		// Create an anonymous function for each goroutine that
		// generates a random number and sends it on the channel.
		go func() {
			ch <- rand.Intn(200)
		}()
	}

	// Create a variable to be used to track received messages.
	// Set the value to the number of goroutines created.
	var numbers []int

	// Iterate receiving each value until they are all received.
	// Store them in a slice of ints
	for emps > 0 {
		p := <-ch
		numbers = append(numbers, p)
		emps--
	}

	// Print the values in our slice
	fmt.Println(numbers)
}
