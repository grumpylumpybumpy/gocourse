package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	message := "Hello, \nGo!"  // \n on a new line
	message1 := "Hello, \tGo!" // \t tab
	message2 := "Hello, \rGo!" // \r character return not a new line
	// it generates Hello then it goes to the very begining and replace Hel with Go!
	message3 := "Hello, Go!"
	rawMessage := `Hello\nGo` // back tic button left of number 1

	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)

	// strings are made up of Rune

	fmt.Println("Length of message variable is", len(message))
	fmt.Println("Length of message3 variable is", len(message3))
	fmt.Println("Length of rawMessage variable is", len(rawMessage))

	fmt.Println("The first character in message var is", message[0]) //ASCII
	greeting := "Hello "
	name := "Alice"
	fmt.Println(greeting + name)

	str1 := "Apple"          // A has an ASCII value of 65
	str := "apple"           // A has an ASCII value of 97
	str2 := "banana"         // b has an ASCII value of 98
	str3 := "app"            // a has an ASCII value of 97
	fmt.Println(str1 < str2) // lexical comparison basically compare length and aphabetical
	fmt.Println(str3 < str1) // compare lengths
	fmt.Println(str > str1)  // lower case vs upper case
	fmt.Println(str > str3)

	// for i, char := range message {
	// 	fmt.Printf("Character at index %d is %c\n", i, char)
	// 	fmt.Printf("%x\n", char) // spits out ASCII value for each character
	// }

	// counts number of characters
	fmt.Println("Rune count:", utf8.RuneCountInString(greeting))

	// strings are immutable cannot add remove etc... must create new string
	greetingWithName := greeting + name
	fmt.Println(greetingWithName)

	var ch rune = 'a' // single quotes are used for runes

	fmt.Println(ch)

	// convert rune to string
	cstr := string(ch)
	fmt.Println(cstr)

	fmt.Printf("Type of cstr is %T\n", cstr)

	// unicode support
	const NIHONGO = "元気ですか？" // Japanese text
	fmt.Println(NIHONGO)

	jhello := "こんにちは" // Hello in Japanese
	for _, runeValue := range jhello {
		fmt.Printf("%c\n", runeValue)
		fmt.Printf("%v\n", runeValue) //unicode value int32
	}

	r := '😊'
	fmt.Println(r)        // print the value
	fmt.Printf("%v\n", r) // print the value as well
	fmt.Printf("%c\n", r) // print the emoji

}
