package main

import (
	"errors"
	"fmt"
)

func main() {

	// func functionName(parameter1 type1, parameter2 type2,...)(returnType1, returnType2,...){
	// 	code block
	// 	return returnValue1, returnValue2,...
	// }

	q, r := divide(10, 3)
	fmt.Printf("Quotient: %v. Remainder: %v\n", q, r)
	// %d is integer %v is any value

	// Error case
	// result, err := compare(3, 2)
	// result, err := compare(3, 4)

	result, err := compare(3, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	// // Testing
	// result, err := compare(3, 3)
	// if err != true {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println(result)
	// }

}

// func divide(a, b int) (int, int) {
// 	quotient := a / b
// 	remainder := a % b
// 	return quotient, remainder
// }

// Another way to express
func divide(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}

func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if b > a {
		return "b is greater than a", nil
	} else {
		return "", errors.New("Unable to compare which is greater")
	}
}

// func compare(a, b int) (string, bool) {
// 	if a > b {
// 		return "a is greater than b", true
// 	} else if b > a {
// 		return "b is greater than a", true
// 	} else {
// 		return "", false
// 	}
// }
