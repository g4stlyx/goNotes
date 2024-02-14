package main

import "fmt"

func defer2(num int) string {
	defer deferFunc()

	if num%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}

func deferFunc() {
	fmt.Println("this will work although there is return up in the function 'defer2' ")
}
