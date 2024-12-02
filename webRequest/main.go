package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Learing web request")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/7") //respone and error
	if err != nil {
		fmt.Println("ERROR", err)
	}
	defer func(Body io.ReadCloser) { //for resource handling
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	fmt.Printf("Type of res: %T\n", res) // *http.Response
	// reading the response

	data, err := ioutil.ReadAll((res.Body))
	if err != nil {
		fmt.Println("ERROR")
	}
	fmt.Println("responese:", string(data)) //as data will be shown in array or bytes we want it in string

}
