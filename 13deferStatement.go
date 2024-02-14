package main

import "fmt"

func a() {
	fmt.Println("the func a has been called")
}

func c() {
	fmt.Println("the func c has been called")
}

func d() {
	fmt.Println("the func d has been called")
}

func b() {
	fmt.Println("the func b has been called")
	// works last regardless of where it's called = even it's in the top of the b(), it will be called after everything else is done in b
	defer a() // "defer" guarantees that func "a()" will be called whatever happens above-errors, "break" etc-.
	defer c()
	defer d() // working order => b, d, c, a
}
