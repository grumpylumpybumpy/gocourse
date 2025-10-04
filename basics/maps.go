package main

import (
	"fmt"
	"maps"
)

func main() {

	// var mapVariable map[keyType]valueType
	// mapVariable = make(map[keyType]valueType)

	// using a Map Literal
	// mapVariable = map[keyType]valueType{
	// 	key1: value1,
	// 	key2: value2
	// }

	myMap := make(map[string]int)
	fmt.Println(myMap)

	myMap["key1"] = 9
	myMap["code"] = 18

	fmt.Println(myMap)
	fmt.Println(myMap["key1"])
	fmt.Println(myMap["key"]) // there is no associated value so the returned value is 0

	// change any value in the map just replace it with new value
	myMap["code"] = 21
	fmt.Println(myMap)

	// delete any value just use the delete function and mention the string associated
	delete(myMap, "key1")
	fmt.Println(myMap)

	myMap["key1"] = 9
	myMap["key2"] = 10
	myMap["key3"] = 11
	fmt.Println(myMap)

	// clear the map completely of all values
	// clear(myMap)
	// fmt.Println(myMap)

	// check if there is a associated value to the key
	value, unknownvalue := myMap["key1"]
	fmt.Println(value)
	fmt.Println(unknownvalue)

	_, unknownvalue1 := myMap["key2"]
	fmt.Println(unknownvalue1)

	// convention
	// _, ok := myMap["key2"]
	// fmt.Println(ok)

	_, ok := myMap["key1"]
	if ok {
		fmt.Println("A value exists with key1")
	} else {
		fmt.Println("No value of a map")
	}

	myMap2 := map[string]int{"a": 1, "b": 2}
	fmt.Println(myMap2)

	myMap3 := map[string]int{"a": 1, "b": 2}

	if maps.Equal(myMap, myMap2) {
		fmt.Println("myMap and myMap2 are equal")
	}

	if maps.Equal(myMap3, myMap2) {
		fmt.Println("myMap and myMap2 are equal")
	}

	for k, v := range myMap3 {
		fmt.Println(k, v)
	}

	for _, v := range myMap3 {
		fmt.Println(v)
	}

	// zero value of a map is nil

	var myMap4 map[string]string

	if myMap4 == nil {
		fmt.Println("The map is initialized to nil value.")
	} else {
		fmt.Println("The map is not initialized to nil value")
	}

	val := myMap["key"]
	fmt.Println(val)

	// myMap4["key"] = "Value"
	// fmt.Println(myMap4)

	// to enter a new value you must use make function
	myMap4 = make(map[string]string)
	myMap4["key"] = "Value"
	fmt.Println(myMap4)

	// the length is counting the numher of keys
	fmt.Println("myMap4 length is", len(myMap4))
	fmt.Println("myMap length is", len(myMap))

	myMap5 := make(map[string]map[string]string)
	myMap5["map1"] = myMap4
	fmt.Println(myMap5)
}
