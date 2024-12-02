package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple, orange, banana"
	// we want comma seperated value in slices
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	str := "one two three two one two five"
	// want to count the occurrence of two

	count := strings.Count(str, "two")
	fmt.Println(count)

	// explore other func like TrimSpaces,  removes space from staring or ending not in between
	// Join for adding

	a := " Hello   go   lang"
	trimmed := strings.TrimSpace(a)
	fmt.Println(trimmed)

	b := "Good"
	c := "Morning"
	res := strings.Join([]string{b, c}, " ")
	fmt.Println(res)

}
