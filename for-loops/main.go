package main

import "fmt"

func main() {
	//there is no while, or do-while loop in go, only for is there.

	//for i := 0; i < 10; i++ {
	//	fmt.Println("Numbers is :", i)
	//
	//}
	//// FOR INFINITE LOOP
	//counter := 0
	//for {
	//	fmt.Println("Infinite Loop")
	//	counter++
	//
	//	if counter == 20 {
	//		break
	//	}
	//}
	//
	////range keyword. for looping and traversing
	nums := []int{9, 5, 4, 33, 23, 11, 53}
	for index, value := range nums {
		fmt.Printf("Index of numbers is %d and value is %d\n", index, value)
	}
	//
	//// we have string data suppose
	//data := "Hello Shriya"
	//for index, value := range data {
	//	fmt.Printf("Index: %d, Value: %c\n", index, value)
	////}

}
