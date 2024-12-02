package main

import "fmt"

func main() {
	age := 25
	name := "Alice"
	height := 161.99908

	fmt.Println("Age of ", name, "is ", age, ". Her Height is", height, "centimeters.")

	//Printf THE FORMATER FUNCTION

	// %d is integer
	fmt.Printf("Age is %d\n", age)
	// here \n is required for  moving to next line

	// %f or %.3f (provide 3 digits after decimal)
	fmt.Printf("Height is %.3f\n", height)

	//for knowing datatype %T is used
	fmt.Printf("Type of name is %T\n", name)

	// %s is for string
	fmt.Printf("name is %s\n", name)

	//all the above can be written in a single line it takes the first var for reference
	fmt.Printf("Name: %s, Age: %d, Height: %.2f", name, age, height)
}
