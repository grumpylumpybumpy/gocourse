package baiscs

import "fmt"

func main() {

	//	// for as while with break
	// sum := 0
	// for {
	// 	sum += 10 // add to sum a numnber of increment
	// 	fmt.Println("Sum:", sum)
	// 	if sum >= 50 {
	// 		break
	// 	}
	// }

	num := 1
	for num <= 10 {
		if num%2 == 0 {
			num++ // ++ increment operator increases by 1 and -- decreases by 1
			continue
		}
		fmt.Println("Odd Number:", num)
		num++
	}

	// ctrl + c will break terminal loop if you are stuck in an infinite loop
}
