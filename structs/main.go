package main

import (
	"fmt"
	"math"
	"time"
)

type Employee struct {
	ID   int
	Name string
	DoB  time.Time
}

type Point struct {
	x, y float64
}

type Path []Point

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func (p *Point) ScaleBy(factor float64) {
	p.x *= factor
	p.y *= factor
}

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q)) // "5", function call
	fmt.Println(p.Distance(q))  // "5", method call
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance()) // "12"

	e1 := Employee{1, "z3", time.Now()}
	e2 := Employee{2, "l4", time.Now()}
	fmt.Printf("e1: %+v\n", e1)
	fmt.Printf("e2.birth: %+v\n", e2.DoB)
}
