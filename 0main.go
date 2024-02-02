package main

import "fmt"

func main() {
	//basics() // $go run 0main.go 1basics.go

	// --------

	// conditionals()

	// --------

	// loops()

	// --------

	// arrays()

	// --------

	//slices() // arraylist karşılığı
	//slices2()

	// --------

	// sum := functions1(3, 5)
	// fmt.Println("", sum)

	// sum2, sub, mltp, div := functions2(3, 5) // returns multiple values
	// sum2, sub, mltp, _ = functions2(4, 2)    // if u dont wanna use a value returned, use _
	// fmt.Println(sum2, sub, mltp, div)

	sum3 := functions3(4, 4, 5, 6, 42, 43352) // use variadic funcs if u dont know how many param'll be given
	sum4 := functions3(2, 3)
	nums := []int{1, 2, 3, 4}
	sum5 := functions3(nums...) // using arrays with variadic functions

	fmt.Println(sum3, sum4, sum5)

	// --------

}
