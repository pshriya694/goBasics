package main

import "fmt"

// defer is used to change execution process -- defer is used to execute the code just before main
func main() {
	fmt.Println("Statement 1")
	defer fmt.Println("Statement 2")
	fmt.Println("Statement 3")

	defer fmt.Println("Statement 4")
	defer fmt.Println("Statement 5")

	//works in LIFO order
	//stacked on the list and when func exits, the statements  are executed in reverse order
}
