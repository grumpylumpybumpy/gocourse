package main

import "fmt"

func main() {
	// panic(interface{})

	// Example of a valid input
	process(10)

	// Example of an invalid input
	process(-3)

}

func process(input int) {

	// Once all defer statements are executed then the panic runs and exit the program
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")

	if input < 0 {
		fmt.Println("Before Panic")
		panic("input must be a non-negative number")
		// NOTE once panic is reached then the rest function exits the function

		// fmt.Println("After Panic")
	}
	fmt.Println("Processing input:", input)
}
