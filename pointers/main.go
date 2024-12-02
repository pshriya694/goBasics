package main

import "fmt"

// pointer is also like a variable, stores address of the data.
// help in memory allocation directly.
func main() {
	var num int
	num = 2

	var ptr *int //pointer which stores data is of int type -- data is of int type

	ptr = &num
	fmt.Println("Num has value:", num, "Has address of ", ptr)

	//// another way of ini
	x := 89
	pt := &x

	fmt.Println(" x has value of:", x, "which address of :", pt)
	// *ptr is  equal to nums //THE DATA
	//*pt is  == x //THE DATA

	//modifying the vale
	value := 10
	modifyValueByReference(&value) // address passed
	fmt.Println("Modified Value is:", value)
	fmt.Println("Address of value is:", &value)

	//for null pointer
	var nullPtr *int
	if nullPtr == nil {
	}
	fmt.Println("Null Pointer:", nullPtr)

}

func modifyValueByReference(a *int) {
	*a = *a * 20

}
