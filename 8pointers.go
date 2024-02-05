package main

import "fmt"

func pointers(num *int) { //* to take the parameter directly from the memory, not as a value.
	//* by this way, the variable given as a paramater changes according to the process done in function.
	*num += 1 //* given number will increase by 1 in main too
	fmt.Println("Address of the given parameter:", num)
	fmt.Println("Number in the function:", *num)
}

//* a func taking array as a parameter would change the original array too. arrays dont work with values, they work w/ references instead.
//* this is same for array based structures like maps, slices etc.
