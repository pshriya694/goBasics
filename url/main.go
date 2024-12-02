package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Parse a URL string into a URL object
	//	myURL :=
	//		"https://example.com:8080/path/to/resource?key1=value1&key2=value2"
	//	parsedURL, err := url.Parse(myURL)
	//	if err != nil {
	//		fmt.Println("Error parsing URL:", err)
	//		return
	//	}
	//	// Accessing URL components
	//	fmt.Println("Scheme:", parsedURL.Scheme)
	//	fmt.Println("Host:", parsedURL.Host)
	//	fmt.Println("Path:", parsedURL.Path)
	//	fmt.Println("RawQuery:", parsedURL.RawQuery)
	//
	//	//Modifying URL components
	//	parsedURL.Scheme = "http"
	//	parsedURL.Host = "newhost.com"
	//	parsedURL.Path = "/newpath"
	//	parsedURL.RawQuery = "key=newvalue"
	//	// Constructing a URL string from a URL object
	//	newURL := parsedURL.String()
	//	fmt.Println("Modified URL:", newURL)

	//encoding
	Person := Person{Name: "John", Age: 30}
	jsonData, err := json.Marshal(Person)
	if err != nil {
		fmt.Println("Error marshalling the data")
		return
	}
	fmt.Println("JSON Data:", string(jsonData))
	// decoding
	newperson := Person
	err = json.Unmarshal(jsonData, &newperson)
	if err != nil {
		fmt.Println("Error unmarshalling the data")
		return

	}
	fmt.Println("New Person:", newperson)

}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
