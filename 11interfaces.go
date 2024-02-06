package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func interfaces1(s shape) { // takes anything which includes area() func inside, bcz shape is created that way
	fmt.Println("The area of the shape:", s.area())
}
