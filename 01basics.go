package main

import "fmt"

func basics() {
	// go'da oluşturduğun her değişkeni, import ettiğin her paketi kullanmak zorundasın
	var str string = "this is a string"
	var num int = 33
	var num2 float64 = 2.0
	num3 := 150
	num4 := 150.2
	isEqual := false
	isEqual = num == num3

	fmt.Println(str)
	fmt.Printf("num: %v, num2: %v, num3: %v, num4: %v \n", num, num2, num3, num4)
	fmt.Print(isEqual)
}
