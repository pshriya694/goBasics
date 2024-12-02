package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("What is your name?")
	var name string
	fmt.Scan(&name)
	//fmt.Println("Hello", name)shriya
	// but can't write the whole name using scan. It read till only white space, hence it breaks.

	//To read the whole line use Bufio Package.

	fmt.Println("What is your favourite food?")

	reader := bufio.NewReader(os.Stdin) //input through os is read by reader.
	food, _ := reader.ReadString('\n')  //see why _ is used!!
	fmt.Println("Good Choice for", food)
	// fmt.Println("Good Choice for", food, "is very tasty") // getting printed on next line see why and correct it.
}
