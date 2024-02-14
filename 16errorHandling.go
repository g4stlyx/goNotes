package main

import (
	"fmt"
	"os"
)

func errorHandling1() {
	f, err := os.Open("a.txt") // success: returns the file, failure: returns the error

	if err != nil {
		// type assertion - tip doğrulama
		if pErr, ok := err.(*os.PathError); ok { // if err == pathError; pErr=err,ok=true
			fmt.Println("file not found: ", pErr.Path)
			return
		} else {
			fmt.Println("another error: ", err)
			return
		}
	} else {
		fmt.Println(f.Name())
	}
}
