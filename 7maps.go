package main

import "fmt"

func maps() { // key-value
	dict := make(map[string]string)
	dict["table"] = "masa"
	dict["book"] = "kitap"
	dict["pen"] = "kalem"

	dict2 := map[string]string{"glass": "bardak", "microphone": "mikrofon"}
	fmt.Println(dict2)

	fmt.Println(dict)
	fmt.Println(dict["table"])
	fmt.Println(len(dict))

	delete(dict, "pen")
	fmt.Println(dict)

	valueOfBook, doesBookExist := dict["book"]
	fmt.Println(valueOfBook, doesBookExist)
}

func forRange() {
	cities := []string{"Ankara", "İstanbul", "İzmir"}

	for i, city := range cities { // If u dont need indices use "_"
		fmt.Print(i, ": ")
		fmt.Println(city)
	}
}

func forRange2() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0

	for _, num := range nums {
		if num%2 != 0 {
			sum += num
		}
	}
	fmt.Println(sum)
}

func forRange3() {
	dict := map[string]string{"book": "kitap", "table": "masa"}

	for k, v := range dict {
		fmt.Println(k, ":", v)
	}
}
