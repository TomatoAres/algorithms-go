package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	Length float64
	Width  float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r Rect) Area() float64 {
	return r.Length * r.Width
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

// TotalPerimeter calculate the total of many shapes perimeter
// ... can pass many parameters
func TotalPerimeter(shape ...Shape) (perimeters float64) {
	for _, s := range shape {
		perimeters += s.Perimeter()
	}

	return perimeters
}

func TotalArea(shape ...Shape) (areas float64) {
	for _, s := range shape {
		areas += s.Area()
	}

	return areas
}

func main() {
	circle := Circle{1}
	rect := Rect{
		Length: 1,
		Width:  2,
	}

	var shapes []Shape = []Shape{circle, rect}

	for _, v := range shapes {
		fmt.Println("Area=", v.Area())
		fmt.Println("Perimeter==", v.Perimeter())
		fmt.Println()
	}

	fmt.Println("totalPerimeter=", TotalPerimeter(circle, rect))
	fmt.Println("TotalArea=", TotalArea(circle, rect))
}
