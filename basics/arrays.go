package main

import "fmt"

func main() {

	// // var arrayName [size]elementType
	// var numbers [5]int //array of 0s
	// fmt.Println(numbers)

	// numbers[4] = 20 //array index starts at 0 so its the 5th column
	// fmt.Println(numbers)

	// numbers[0] = 9
	// fmt.Println(numbers)

	// fruits := [4]string{"Apple", "Banana", "Orange", "Grapes"}
	// fmt.Println("Fruits array:", fruits)

	// fmt.Println("Third element:", fruits[2])

	// originalArray := [3]int{1, 2, 3}
	// copiedArray := originalArray

	// copiedArray[0] = 100

	// fmt.Println("Orginal array:", originalArray)
	// fmt.Println("Copied array:", copiedArray)

	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println("Element at index,", i, ":", numbers[i])
	// }

	// // for index, value := range numbers {
	// // 	fmt.Printf("Index: %d, Value: %d\n", index, value)
	// // }

	// // for i, v := range numbers { // i replaces index v replaces value _ do not want to return a value
	// // 	fmt.Printf("Index: %d, Value: %d\n", i, v) // use printf instead of println because we are printing placeholders
	// // }

	// // underscore is blank identifier, used to store unsued values
	// for _, v := range numbers { // _ discard the value blank identifier
	// 	fmt.Printf("Value: %d\n", v) // use printf instead of println
	// }

	// // a, _ := someFunction() //
	// // fmt.Println(a)
	// // // fmt.Println(b)

	// // b := 2
	// // _ = b // demonstration only this doesn't make sense

	// fmt.Println("The length of the numbers array is ", len(numbers))

	// // Comparing Arrays
	// array1 := [3]int{1, 2, 3}
	// array2 := [3]int{1, 2, 3}

	// fmt.Println("Array1 is equal to Array2: ", array1 == array2)

	// var matrix [3][3]int = [3][3]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }

	// fmt.Println(matrix)

	originalArray := [3]int{1, 2, 3}
	var copiedArray *[3]int // just an array of 3 elements

	// copiedArray is a pointer, address of orginalArray
	copiedArray = &originalArray // copiedArray is now the actual orginalArray
	copiedArray[0] = 100

	fmt.Println("Orginal array:", originalArray)
	fmt.Println("Copied array:", *copiedArray) //the * stops the & from printing

}

func someFunction() (int, int) {
	return 1, 2
}
