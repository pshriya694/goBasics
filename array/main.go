package main

import "fmt"

func main() {
	fmt.Println("Learning Array in GO lang")

	var name [5]string
	name[0] = "abc"
	name[2] = "xyz"

	//fmt.Println("Here is content", name)
	//
	//// another way of initializing and declaring array
	//var numbers = [5]int{3, 6, 9, 11, 13}
	//fmt.Println("Number are:", numbers)
	//
	////getting length of array
	//fmt.Println("length of number array  is:", len(numbers))
	//
	////acess
	//fmt.Println("value at 2nd index:", name[2])

	// array initialized set to 0.
	var price [5]int
	fmt.Println("Price is:", price) // [0,0,0,0,0]

	var prices [5]bool //[false,false,false,false,false]
	fmt.Println("Price is:", prices)

	var pricess [5]string //empty array
	fmt.Println("Price is:", pricess)
	//to see how it is init, we use %q stands for quoted string.

	fmt.Printf("Price Array is %q\n", pricess) //Price Array is ["" "" "" "" ""] //quotation string is used

	//for pointers and complex type it's nil
	//will be learned furthur.

	//array has a prob of size, we need to define it. but not so with slices.
}
