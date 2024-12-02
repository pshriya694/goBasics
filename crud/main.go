package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   // No `json` tag specified
}

func PerformGetRequest() {
	const myurl = "https://jsonplaceholder.typicode.com/todos/1"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		fmt.Println("Failed to fetch data:", response.Status)
		return
	}
	var todo Todo
	err = json.NewDecoder(response.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}
	fmt.Println("Todo ID:", todo.ID)
	fmt.Println("Todo Title:", todo.Title)
	fmt.Println("Todo Completed:", todo.Completed)
}
func main() {
	PerformGetRequest()
}
