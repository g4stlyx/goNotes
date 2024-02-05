package main

import "fmt"

func structs() {
	fmt.Println(product{"computer", 500, "asus"})
	fmt.Println(product{name: "another computer"}) // other places take default values(for int its 0, for string its "")
}

type product struct { //* like classes/models
	name      string
	unitPrice float64
	brand     string
}

//---

func structs2() {
	c := customer{"sefa", "a", 33}
	c.save()
	c.update()
}

type customer struct {
	firstName string
	lastName  string
	age       int
}

func (c customer) save() {
	fmt.Println("saved:", c.firstName)
}

func (c customer) update() {
	fmt.Println("updated:", c.firstName)
}
