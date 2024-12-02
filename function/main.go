package main

import "fmt"

func simpleFunction() { // this curly bracket must start from same line as the func defined otherwise it give error.
	fmt.Println("Returning a Statement")
}
func add(a, b int) int { // output type (int) input type (a,b int), also if single return type is there then no need for bracket
	return a + b
}
func sub(a, b int) int {
	return a - b
}

// //func mul(a int, b float64) (product float64) {
// //	product = a * b
// //	return
// //}
// // check how to pass multi data type
func main() {
	fmt.Println("Learning Functions")
	//simpleFunction()
	//
	//ans := add(6, 4) //passing parameters
	//fmt.Println("Sum of a and b is: ", ans)

	a := sub(5, 3)
	fmt.Println("Subtraction of a and b is: ", a)

	//mul(5, 9.8)
	//fmt.Println("Product of a and b is", product)
}
