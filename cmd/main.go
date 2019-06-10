package main

import (
	"fmt"

	"github.com/rmartinus/practice-my-go/pkg/shape"
)

func main() {
	s := shape.Square{5}
	fmt.Println("Square:", s)
	fmt.Println("Area of square:", s.Area())
	s.ScaleBy(2)
	fmt.Println("Square scaled by 2:", s)

	r := shape.Rectangle{2, 7}
	fmt.Println("Rectangle:", r)
	fmt.Println("Area of rectangle:", r.Area())
	r.ScaleBy(3)
	fmt.Println("Rectangle scaled by 3:", r)

	c := shape.Cube{3}
	fmt.Println("Cube:", c)
	fmt.Println("Area of cube:", c.Area())
	fmt.Println("Volume of cube:", c.Volume())
	c.ScaleBy(4)
	fmt.Println("Cube scaled by 4:", c)
}
