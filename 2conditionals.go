package main

import "fmt"

func conditionals() {
	var balance float64 = 1200.24
	var moneyToBeWithdrawn float64 = 900

	if moneyToBeWithdrawn > balance {
		fmt.Print("Not enough money!")
	} else {
		fmt.Printf("%v TL is successfully withdrawn.", moneyToBeWithdrawn)
	}

}
