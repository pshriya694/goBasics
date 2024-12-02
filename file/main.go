package main

import (
	"fmt"
	"io"
	"os"
)

// involves file handling with os ans io/ioutil package
func main() {
	file, err := os.Create("example.text")
	if err != nil {
		fmt.Println("ERROR")
	} //returning type and error
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file) //otherwise resource may leak

	//content for the file
	content := "Hello This is initial content"
	_, err = io.WriteString(file, content+"\n") // for adding the content
	content = "Hello This is second content"
	_, err = io.WriteString(file, content+"\n") // with a new line
	if err != nil {
		fmt.Println("ERR while writing the file", err)

	}

	fmt.Println("CREATED A FILE")

	//see 24 video.  use package

	// content, err = ioutil.ReadFile("emample.text")

	// when have time look in details for file handling and time package.
	//-------------------------------------------//
	file, err = os.Open("example.text")
	if err != nil {
		fmt.Println("ERROR")
		return
	}
	defer file.Close()

	//creating buffer
	buffer := make([]byte, 1024)

	//reading the file

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break //END OF FILE REACH

		}
		if err != nil {
			fmt.Println("ERROR WHILE READING FILE", err)
			return
		}
		//processing reading content
		fmt.Print(string(buffer[:n]))

	}
}
