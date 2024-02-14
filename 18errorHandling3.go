package main

import "fmt"

type BorderException struct {
	parameter int
	message   string
}

func (b *BorderException) Error() string { // to implement Error() interface in BorderException
	return fmt.Sprintf("%d----%s", b.parameter, b.message)
}

func guess2(guess_num int) (string, error) {
	num := 15
	if guess_num < 1 || guess_num > 100 {
		return "", &BorderException{parameter: guess_num, message: "Number must be between 1 and 100"}
	} else if guess_num < num {
		return "go up", nil
	} else if guess_num > num {
		return "go down", nil
	} else {
		return "u r goddamn right", nil
	}
}

func errorHandling3() {
	msg1, _ := guess2(50)
	fmt.Println(msg1)
	fmt.Println(guess2(101))
	fmt.Println(guess2(-10))
}
