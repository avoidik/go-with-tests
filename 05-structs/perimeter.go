package perimeter

import "math"

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
}

func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Perimeter() float64 {
	return c.Radius * math.Pi
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}

func (t Triangle) Area() float64 {
	s1 := (t.SideA + t.SideB + t.SideC) / 2
	return math.Sqrt(s1 * (s1 - t.SideA) * (s1 - t.SideB) * (s1 - t.SideC))
}
