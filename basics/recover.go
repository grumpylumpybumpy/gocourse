package main

import "fmt"

func main() {

	process()
	fmt.Println("Returned from process")
}

func process() {
	defer func() {
		// this checks if there was a panic
		// if r := recover(); r != nil {
		r := recover()
		if r != nil {
			fmt.Println("Recover:", r)
		}
	}()

	fmt.Println("Start Process")
	panic("Something went wrong")
	fmt.Println("End Process")
}
