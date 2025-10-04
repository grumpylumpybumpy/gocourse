package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	str := "Hello Go!"

	fmt.Println(len(str))

	str1 := "Hello"
	str2 := "World!"

	result := str1 + " " + str2
	fmt.Println(result)
	fmt.Println(str[0])

	fmt.Println(str)
	fmt.Println(str[1:7])

	// String Conversion
	num := 18
	str3 := strconv.Itoa(num)
	// fmt.Println(len(num)) // because num is not a string
	fmt.Println(len(str3))

	// String splitting
	fruits := "apple, orange, banana"
	fruits1 := "apple-orange-banana"
	parts := strings.Split(fruits, ",")
	parts1 := strings.Split(fruits1, "-")
	fmt.Println(parts)
	fmt.Println(parts1)

	countries := []string{"Germany", "France", "Italy"}
	joined := strings.Join(countries, ", ")
	fmt.Println(joined)

	fmt.Println(strings.Contains(str, "Go?"))

	replaced := strings.Replace(str, "Go", "Universe", 1)
	fmt.Println(replaced)

	strwspace := " Hello Everyone! "
	fmt.Println(strwspace)
	fmt.Println(strings.TrimSpace(strwspace))

	fmt.Println(strings.ToLower(strwspace))
	fmt.Println(strings.ToUpper(strwspace))

	fmt.Println(strings.Repeat("foo ", 3))

	fmt.Println(strings.Count("Hello", "l"))
	fmt.Println(strings.HasPrefix("Hello", "he"))
	fmt.Println(strings.HasSuffix("Hello", "lo"))
	fmt.Println(strings.HasSuffix("Hello", "la"))

	str5 := "Hello, 123 Go!"
	re := regexp.MustCompile()
}
