package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Todo struct {
	// to do it in an easier way, use online tools like "json-to-go"
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func restful() {
	// get
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1") // to make requests in go
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	bodyBytes, _ := io.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString) // string

	var todo Todo
	json.Unmarshal(bodyBytes, &todo)
	fmt.Println(todo) // struct/class, mostly this is used. by this way, you can take whatever you want from the response

}

func restful2() { //post
	todo := Todo{1, 2, "this is smh to do", false}
	jsonTodo, jsonErr := json.Marshal(todo)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	response, resErr := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json;charset=utf-8", bytes.NewBuffer(jsonTodo))
	if resErr != nil {
		fmt.Println(resErr)
	}
	defer response.Body.Close()

	bodyBytes, _ := io.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString) // string

	var todo2 Todo
	json.Unmarshal(bodyBytes, &todo2)
	fmt.Println(todo)

}
