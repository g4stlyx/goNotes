package main

func main() {
	//* basics() // $go run 0main.go 1basics.go
	//* conditionals()
	//* loops()
	//* arrays()

	//* slices -arraylists-
	//slices() // arraylist karşılığı
	//slices2()

	//* functions
	// sum := functions1(3, 5)
	// fmt.Println("", sum)

	// sum2, sub, mltp, div := functions2(3, 5) // returns multiple values
	// sum2, sub, mltp, _ = functions2(4, 2)    // if u dont wanna use a value returned, use _
	// fmt.Println(sum2, sub, mltp, div)

	// sum3 := functions3(4, 4, 5, 6, 42, 43352) // use variadic funcs if u dont know how many param'll be given
	// sum4 := functions3(2, 3)
	// nums := []int{1, 2, 3, 4}
	// sum5 := functions3(nums...) // using arrays with variadic functions

	// fmt.Println(sum3, sum4, sum5)

	//* maps and for range(forEach)
	//maps()
	//forRange()
	//forRange2()
	//forRange3()

	//* pointers
	// number := 33
	// pointers(&number)
	// fmt.Println("Number in main:", number)

	//* structs
	//structs2()

	//* go routines and channels, async programming
	// channels are used to control/see when the processes are done
	// oddNumberCn := make(chan int)
	// evenNumberCn := make(chan int)
	// go oddNumbers(oddNumberCn)   // "go" is used to make them work async
	// go evenNumbers(evenNumberCn) // "go" is used to make them work async

	// oddNumberSum, evenNumberSum := <-oddNumberCn, <-evenNumberCn // getting int values from int channels

	// product := oddNumberSum * evenNumberSum
	// fmt.Println("Their product:", product)

	//* interfaces
	// c := circle{radius: 5}
	// r := rectangle{width: 10, height: 6}
	// interfaces1(c) // takes a shape as parameter
	// interfaces1(r)

	//interfaces2()

	//* defer
	// when we want some funcs to be called inside of a func.
	// generally used for logging purposes or e.g closing the db connection
	// b()

	// var result = defer2(14)
	// fmt.Println(result)
	// defer3()

	//* debugging
	// errorHandling1()
	// errorHandling2()
	// errorHandling3()

	//* string functions
	// stringFunctions()

	//* restful
	// restful()
	restful2()
}
