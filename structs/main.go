package main

import (
	"fmt"
	"github.com/compasty/hello-go/structs/bitmap"
	"image/color"
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

type ColoredPoint struct {
	Point
	Color color.RGBA
}

type Path []Point

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

// same thing, but as a method of the *Point type
func (p *Point) Distance(q *Point) float64 {
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
			sum += path[i-1].Distance(&path[i])
		}
	}
	return sum
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q)) // "5", function call
	fmt.Println(p.Distance(&q)) // "5", method call
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

	cp1 := ColoredPoint{Point{1, 1}, color.RGBA{255, 0, 0, 255}}
	cp2 := ColoredPoint{Point{2, 3}, color.RGBA{0, 255, 0, 255}}
	fmt.Println(cp1.Distance(&cp2.Point))

	var x, y bitmap.Bitmap
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String())           // "{1 9 42 144}"
	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	fmt.Println(x.Len())              // 4
	x.Remove(9)
	fmt.Println(x.String())
}
