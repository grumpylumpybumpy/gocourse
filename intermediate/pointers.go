package main

import "fmt"

func main() {
	// pointer store memory address
	// var ptr *int

	var ptr *int
	var a int = 10
	ptr = &a // referencing

	fmt.Println(a)
	fmt.Println(ptr)

	// value in a is stored at the address we printed the address
	// fmt.Println(&a)

	// use * twice is de-reference the pointer
	// fmt.Println(*ptr) // de-reference a pointer

	// if ptr == nil {
	// 	fmt.Println("Pointer is nil")
	// }

	modifyValue(ptr)
	fmt.Println(a)

}

func modifyValue(ptr *int) {
	*ptr++
}
