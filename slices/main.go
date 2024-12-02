package main

import "fmt"

// slices are flexible and dynamic DS.
//powerful alternative to array.

func main() {
	fmt.Println("-----Learning Slices-----")
	numbers := []int{7, 8, 9, 5, 44, 56, 32, 11, 2, 3, 4}
	//fmt.Println(numbers)

	//fmt.Printf("Number has data type of %T\n", numbers)
	//fmt.Println("Length of numbers is", len(numbers))
	//
	////adding numbers to array using append
	//numbers = append(numbers, 99, 76, 46, 12)
	//fmt.Println(numbers)
	//fmt.Println("Length now", len(numbers))
	//fmt.Println("Capacity is", cap(numbers))
	//
	////use of make function
	////takes length and capacity as input to make slices
	////   https://golangforall.com/en/post/golang-slice.html  for slice learning
	//
	////using make func, we can set length and capacity like we want to.

	numbers = make([]int, 3, 5)
	fmt.Println("Length now", len(numbers))
	fmt.Println("Capacity is", cap(numbers))
	fmt.Println(numbers)

	numbers = append(numbers, 4, 9, 99, 00, 76, 88, 87, 32, 87, 456)
	fmt.Println("Length now", len(numbers))
	fmt.Println("Capacity is", cap(numbers))
	fmt.Println(numbers)

	numbers = append(numbers, 3, 44, 22)
	fmt.Println("Length now", len(numbers))
	fmt.Println("Capacity is", cap(numbers))
	fmt.Println(numbers)

	//can't initialize like this  var numbers = []string // have to use make func
	//got empty slice use this
	//stock := []string{} //or use a make func stock := make([]int, 0)
	//fmt.Println(stock)

}
