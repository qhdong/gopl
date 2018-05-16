package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	// var w Wheel
	// w.Circle.Center.X = 20
	var w = Wheel{
		Circle: Circle{
			Point: Point{
				X: 20,
				Y: 30,
			},
			Radius: 5,
		},
		Spokes: 10,
	}
	w.Spokes = 20

	fmt.Printf("w = %+v\n", w)
}
