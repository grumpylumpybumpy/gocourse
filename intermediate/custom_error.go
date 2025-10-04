package main

import (
	"errors"
	"fmt"
)

func main() {

	err := doSomething()
	if err != nil {
		fmt.Print(err)
		// we need to return so we can make sure the code does not continue to run
		return
	}
	fmt.Println("Operation completed successfully")

}

type customError struct {
	code    int
	message string
	er      error // wrap error
}

// Error returns the error message. Implementing Error() method of error interface
//
//	func (e *customError) Error() string {
//		return fmt.Sprintf("Error %d: %s\n", e.code, e.message)
//	}
func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s, %v\n", e.code, e.message, e.er)
}

// Function that return a custom error
// func doSomething() error {
// 	return &customError{
// 		code:    500,
// 		message: "Something went wrong!",
// 	}
// }

func doSomething() error {
	err := doSomethingElse()
	if err != nil {
		return &customError{
			code:    500,
			message: "Something went wrong",
			er:      err,
		}
	}
	return nil
}

func doSomethingElse() error {
	return errors.New("Internal error")
}
