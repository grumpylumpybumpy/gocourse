package main

import "fmt"

func main() {

	// Switch statement (switch case default automatic break after each case)
	// switch expression {
	// case value1:
	// 	// code to be executed if expression equals value1
	// 	fallthrough // this will make the code continue to the next case
	// case value2:
	// 	// code to be executed if expression euqals value2
	// case value3:
	// 	// code to be executed if expression equals value3
	// default:
	// 	// code to be executed if expression does not match any value
	// }

	// Switch statement in other languages (switch case break default)
	// switch expression {
	// case value1:
	// 	// code to be executed if expression equals value1
	// 	break;
	// case value2:
	// 	// code to be executed if expression euqals value2
	// 	break;
	// case value3:
	// 	// code to be executed if expression equals value3
	// 	break;
	// default:
	// 	// code to be executed if expression does not match any value
	// 	break;
	// }

	// fruit := "pinapple"

	// switch fruit {
	// case "apple":
	// 	fmt.Println("It's an apple.")
	// case "banana":
	// 	fmt.Println("It's a banana.")
	// default:
	// 	fmt.Println("Unknown Fruit!")
	// }

	// // Multiple Conditions
	// day := "Monday"
	// switch day { // multiple cases put them together
	// case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
	// 	fmt.Println("It's a weekday.")
	// case "Sunday":
	// 	fmt.Println("It's a weekend.")
	// default:
	// 	fmt.Println("Invalid day.")
	// }

	// // Multiple Expression
	// number := 15
	// switch {
	// case number < 10:
	// 	fmt.Println("Number is less than 10")
	// case number >= 10 && number < 20:
	// 	fmt.Println("Number is betwen 10 and 19")
	// default:
	// 	fmt.Println("Number is 20 or more")
	// }

	// // Fallthrough
	// num := 2
	// switch {
	// case num > 1:
	// 	fmt.Println("Greater than 1")
	// 	fallthrough
	// case num == 2:
	// 	fmt.Println("Number is 2")
	// default:
	// 	fmt.Println("Not Two")
	// }

	// checkType(10)
	// checkType(3.14)
	// checkType("Hello")
	// checkType(true)

	// if else is ok to use if there are only less than 10 checks
	// switch is better used for longer checks cleaner and readable

}

func checkType(x interface{}) { // when datatype not explicitly declared
	switch x.(type) {
	case int:
		fmt.Println("It's an integer")
		// fallthrough is not permitted to use in switch type because it is illogical
	// case int32:
	// 	fmt.Println("it's an integer 32")
	case float64:
		fmt.Println("It's a float")
	case string:
		fmt.Println("It's a string")
	default:
		fmt.Println("Unknown Type")
	}
}
