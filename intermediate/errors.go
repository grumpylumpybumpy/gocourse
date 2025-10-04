package main

import (
	"errors"
	"fmt"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Math Error: square root of a negative number")
	}
	// compute the square root
	return 1, nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("Error: Empty data")
	}
	// Process data
	return nil
}

func main() {

	// result, err := sqrt(16)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(result)

	// result1, err1 := sqrt(-16)
	// if err1 != nil {
	// 	fmt.Println(err1)
	// 	return
	// }
	// fmt.Println(result1)

	// data := []byte{}
	// if err := process(data); err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println("Data Processed Successfully")

	// data := []byte{}
	// // This is a short version of the next two lines
	// // if err := process(data); err != nil {
	// err := process(data)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println("Data Processed Successfully")

	// // Error interface of built in package
	// err1 := eprocess()
	// if err1 != nil {
	// 	fmt.Println(err1)
	// 	return
	// }
	// fmt.Println("Data Processed Successfully")

	// // Built in package no need to import
	// println("")

	err := realData()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data read successfully.")

}

type myError struct {
	message string
}

func (m *myError) Error() string {
	return fmt.Sprintf("Error: %s", m.message)
}

func eprocess() error {
	return &myError{"Custom error message"}
}

func realData() error {
	err := realConfig()
	if err != nil {
		return fmt.Errorf("realData: %w", err)
	}
	return nil
}

func realConfig() error {
	return errors.New("config error")
}
