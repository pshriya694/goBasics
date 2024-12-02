package main

import "fmt"

// maps retrieve data via key, it is unordered
// similar to hashmap and dictionaries
//has key which allow efficient retrieval of data

func main() {
	//name -(string)<-> marks (int)
	studentsGrades := make(map[string]int)
	//data filling

	studentsGrades["Riya"] = 67
	studentsGrades["Nikhil"] = 89
	studentsGrades["Priyanshu"] = 45
	studentsGrades["Deepika"] = 90
	studentsGrades["Rohini"] = 59

	// data retrieval
	fmt.Println("Marks of Deepika:", studentsGrades["Deepika"]) //key must be as it is.
	//value change
	studentsGrades["Deepika"] = 94
	fmt.Println("Marks of Deepika:", studentsGrades["Deepika"])

	//deleting data
	delete(studentsGrades, "Priyanshu")
	fmt.Println("Priyanshu marks:", studentsGrades["Priyanshu"])

	////checking if key is present or not
	////use of boolean
	////var exist bool
	//
	_, exist := studentsGrades["Priyanshu"]
	fmt.Println("Exist:", exist)
	//
	//grades, exist = studentsGrades["Deepika"]
	//fmt.Println("Grade:", grades, "Exist:", exist)
	//
	////in for loop
	//
	//for index, value := range studentsGrades {
	//	fmt.Printf("key is %s and marks is %d\n", index, value)
	//}
	//
	//another way of making a map

	person := map[string]int{
		"Mia":     99,
		"Dave":    98,
		"Charlie": 92,
	}
	for index, value := range person {
		fmt.Printf("-------Key is %s and marks is %d\n", index, value)
	}
}
