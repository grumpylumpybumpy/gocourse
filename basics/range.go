package main

import "fmt"

func main() {

	message := "HeLlo World"

	for i, v := range message {
		// fmt.Println(i, v)
		// %c is the code %d is value
		fmt.Printf("Index: %d, Rune: %c\n", i, v)
	}

	for _, v := range message {
		// fmt.Println(i, v)
		// %c is the code %d is value
		fmt.Printf("Rune: %c\n", v)
	}

}
