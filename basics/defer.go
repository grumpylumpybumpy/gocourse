package main

import "fmt"

func main() {

	// defer statement invokes a function to executes once the surrounding function spitsan output
	process(10)

}

func process(i int) {
	// NOTE the first statement to defer will run last
	defer fmt.Println("Deferred i value:", i)
	defer fmt.Println("First deferred statement executed")
	defer fmt.Println("Second deferred statement executed")
	defer fmt.Println("Third deferred statement executed")
	i++
	fmt.Println("Normal execution statement")
	fmt.Println("Value of i:", i)
}
