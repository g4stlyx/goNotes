package main

import "fmt"

func loops() {
	number := 0
	fmt.Println("Enter a positive integer")
	fmt.Scanln(&number)

	isPrime := true
	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false
		}
	}

	if isPrime {
		fmt.Printf("%v is prime \n", number)
	} else {
		fmt.Printf("%v is not prime \n", number)
	}

}
