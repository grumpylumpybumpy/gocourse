package main

import "fmt"

func main() {

	// if condition {
	// 	block of code
	// }

	// age := 25

	// if age >= 18 {
	// 	fmt.Println("You are an adult.")
	// }

	// if condition {
	// 	block of code
	// } else if {
	// 	block of code
	// } else {
	// 	block of code
	// }

	// temperature := 25
	// if temperature >= 30 {
	// 	fmt.Println("It's hot outside.")
	// } else {
	// 	fmt.Println("It's cool outside.")
	// }

	// score := 95
	// // as soon as the first condition is satisfy then it exits the code
	// if score >= 90 {
	// 	fmt.Println("Grade A")
	// } else if score >= 80 {
	// 	fmt.Println("Grade B")
	// } else if score >= 70 {
	// 	fmt.Println("Grade C")
	// } else {
	// 	fmt.Println("Grade D")
	// }

	// the order of the statements are also important
	// checks all the statements except for the else statement since the else statement belongs with the if
	// if score >= 90 {
	// 	fmt.Println("Grade A")
	// } // note the next if is on a different line
	// if score >= 80 {
	// 	fmt.Println("Grade B")
	// }
	// if score >= 70 {
	// 	fmt.Println("Grade C")
	// } else {
	// 	fmt.Println("Grade D")
	// }

	// if condition1 {
	// 	code block
	// 	if condition2 {
	// 		code block2
	// 	}
	// }

	// num := 6
	// if num%2 == 0 {
	// 	if num%3 == 0 {
	// 		fmt.Println("Numner is divisible by both 2 and 3.")
	// 	} else {
	// 		fmt.Println("Number is divisble by 2 but not 3.")
	// 	}
	// } else {
	// 	fmt.Println("Number is not divisible by 2")
	// }

	// || OR
	// && AND

	if 10%2 == 0 || 5%2 == 0 { // if the first condition is true the second condition is not checked
		fmt.Println("Either 10 or 5 are even.")
	}

	if 10%2 == 0 && 6%2 == 0 { // check both conditons regardless if the first condition is true
		fmt.Println("Both 10 and 6 are even.")
	}
}
