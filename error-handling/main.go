package main

import "fmt"

func divide(a, b float64) float64 {
	//if b == 0 {
	//	return 0, fmt.Errorf("denominator must not be Zero") //two parameter are there
	//}
	return a / b

}
func main() {
	fmt.Println("-----Error Handling in Functions------")
	//ans, _ := divide(10, 0) // in go _ is used as blank identifier. It serves as a write-only variable that allows you to discard values returned by func

	ans := divide(10, 0)
	//if err != nil {
	//	fmt.Println("Error Handling")
	//}

	fmt.Println("Division is:", ans)
}
