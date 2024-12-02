package main

import (
	"fmt"
	"mylearning/myutil"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("--------------------------------------")
	myutil.PrintMessage("HII")
	fmt.Println("--------------------------------------")
	//
	////variables in go
	//var name = "Shriya"
	//fmt.Println(name)
	//
	//var version = 1.0 //version has string, go is smart.
	//fmt.Println(version)
	//
	//var a int = 78 //writing integer is not necessary.
	//fmt.Println(a)
	//
	//var b float64 = 90.55
	//fmt.Println(b)
	//
	//var person = "Riy Dutt"
	//fmt.Println(person)
	//
	//// use of constant
	//
	//const pi = 3.14159
	////const won't let it change again like this.
	//// pi = 22/7 , this will throw an error
	//fmt.Println(pi)
	//
	//var decided bool = false
	//decided = true
	//fmt.Println(decided)
	//
	////making good variables without declaration
	//
	//human := "mahi" // used when we are getting data directly from source.
	//fmt.Println(human)
	//
	////IMPORTANT: CAPITALIZATION OF NAME
	//var Public = "Data is Shared"   // name must start with capital so that it is accessible in all file. //by that i mean package
	//var private = "Data is private" // only accessible in this file, not outside.
	//
	//fmt.Println(Public)
	//fmt.Println(private)

	fmt.Println(myutil.PublicMessage)
	fmt.Println(myutil.PrivateMessage)

	fmt.Println("hello world")

}
