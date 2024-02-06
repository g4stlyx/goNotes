package main

// done by/with routines and channels

func oddNumbers(oddNumbersCn chan int) {
	sum := 0
	for i := 1; i <= 10; i += 2 {
		sum = sum + i
	}

	oddNumbersCn <- sum
}

func evenNumbers(evenNumbersCn chan int) {
	sum := 0
	for i := 0; i <= 10; i += 2 {
		sum = sum + i
	}

	evenNumbersCn <- sum
}
