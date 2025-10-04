package main

import (
	"fmt"
	"os"
)

func main() {
	// exit stops all immediately
	fmt.Println("Starting the main function")

	// Exit with the status code of 1
	os.Exit(1)

	// This will never be executed
	fmt.Println("End of main function")
}
