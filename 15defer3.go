package main

import "fmt"

type product2 struct {
	productName string
	unitPrice   int
}

func (p product2) save() {
	fmt.Println("Prodcut saved: ", p.productName)
}

func defer3() {
	p := product2{productName: "Laptop", unitPrice: 25000}
	defer p.save() // since "p" used here is the last one before defer statement, p.productName = "Laptop"
	p = product2{productName: "Mouse", unitPrice: 550}
	defer p.save() // since "p" used here is the last one before defer statement, p.productName = "Mouse"
	fmt.Println("Process Successful")

}
