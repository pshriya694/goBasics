package main

import "fmt"

// struct is complex DS.
// Has multi data types act as single Data type
type Person struct {
	FirstName string
	Lastname  string
	Age       int
}

func main() {
	//using var keyword
	//var Emp1 Person
	//fmt.Println("Priya Person:", Emp1) //o/p {  0} -- space, space, 0 (default value)
	//fmt.Println("---------------------------------------------")
	//Emp1.FirstName = "Priyanshi"
	//Emp1.Lastname = "Aggrawal"
	//Emp1.Age = 19
	//fmt.Println("Priya Person:", Emp1)
	//
	//// second method: struct literal
	//Emp2 := Person{
	//	FirstName: "Aakash",
	//	Lastname:  "Singh",
	//	Age:       21,
	//}
	//fmt.Println("People:", Emp2)
	//fmt.Println(Emp2.Lastname)
	//
	//// 3rd method: using new keyword (pointer)
	//var Emp3 = new(Person)
	//Emp3.FirstName = "ABC"
	//Emp3.Lastname = "WXY"
	//Emp3.Age = 51
	//fmt.Println("person2:", Emp3) //like a pointer.
	//fmt.Println("person2:", Emp3.FirstName)

	//STRUCT EMBEDDING (embedding within other struct)
	// creating an instance of employee struct
	Emp := Employee{
		Person: Person{
			FirstName: "Joey",
			Lastname:  "Tribbiani",
			Age:       40,
		},
		Address: Address{
			Street:  "34th Street",
			City:    "New York",
			Country: "USA",
		},
		Contact: Contact{
			Phone: "1234567890",
			Email: "abc@gmail.com",
		},
		Position: "Developer",
	}
	fmt.Println("Employee:", Emp)

}

// struct embedding (embedding within other struct)
type Address struct {
	Street  string
	City    string
	Country string
}

type Contact struct {
	Phone string
	Email string
}
type Employee struct {
	Person
	Address
	Contact
	Position string
}

//use for modeling data and organizing data, building relationships
