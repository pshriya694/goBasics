package main

import (
	"fmt"
	"strconv"
)

func main() {
	//num := 67
	//fmt.Printf("Num is of type %T\n", num)
	//
	//data := float64(num) // can also be -- var data float64 = float64(num)
	//fmt.Printf("Data is of type %T\n", data)

	//for integer to string conversion
	num := 123
	str := strconv.Itoa(num) //itoa is string conversion func in integer to string  - integer to a

	fmt.Printf("DATA TYPE FOR NUM IS %T\n", str)

	num_string := "1234"
	num_int, _ := strconv.Atoi(num_string) //ignoring the error field
	num_int = num_int + 990099             //after converting
	fmt.Printf("DATA TYPE FOR NUM IS %d\n", num_int)

}
