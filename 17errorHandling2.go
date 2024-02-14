package main

import (
	"errors"
	"fmt"
)

func guess(guess_num int) (string, error) {
	num := 15
	if guess_num < 1 || guess_num > 100 {
		return "", errors.New("enter a number between 1 and 100")
	} else if guess_num < num {
		return "go up", nil
	} else if guess_num > num {
		return "go down", nil
	} else {
		return "u r goddamn right", nil
	}

}

func errorHandling2() {
	msg1, _ := guess(50)
	fmt.Println(msg1)
	fmt.Println(guess(101))
	fmt.Println(guess(-10))
}
