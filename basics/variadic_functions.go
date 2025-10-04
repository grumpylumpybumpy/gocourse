package main

import "fmt"

func main() {

	// ... Ellipsis indicate function can accept 0 or more arguement of that type
	// func functionName(param1 type1, param2 type2, param3 ... type3) returnType {
	// 	function body
	// }

	// fmt.Println("Sum of 1, 2, 3:", sum(1, 2, 3))

	// statement, total := sum("The sum of 1, 2, 3, 4, 5, 6 is", 1, 2, 3, 4, 5, 6)
	// fmt.Println(statement, total)

	sequence, total := sum(1, 20, 30, 40, 50, 60)
	fmt.Println("Sequence:", sequence, "Total", total)
	sequence2, total2 := sum(2, 40, 36, 40, 50, 60)
	fmt.Println("Sequence:", sequence2, "Total", total2)

	// slice example
	numbers := []int{1, 2, 3, 4, 5, 9}
	// turn the slice into a variadic parameter by adding ...
	sequence3, total3 := sum(3, numbers...)
	fmt.Println("Sequence: ", sequence3, "Total", total3)

}

// func sum(nums ...int) int {
// 	total := 0
// 	for _, v := range nums {
// 		total += v
// 	}
// 	return total
// }

// NOTE variadic parameter must be LAST!

// func sum(returnString string, nums ...int) (string, int) {
// 	total := 0
// 	for _, v := range nums {
// 		total += v
// 	}
// 	return returnString, total
// }

func sum(sequence int, nums ...int) (int, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return sequence, total
}
