package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y int
}

type Lingkaran struct {
	center Titik
	radius int
}

func jarak(p, q Titik) float64 {
	return math.Sqrt(
		math.Pow(float64(p.x-q.x), 2) +
			math.Pow(float64(p.y-q.y), 2))
}

func didalam(c Lingkaran, p Titik) bool {
	return jarak(c.center, p) <= float64(c.radius)
}

func main() {
	var c1, c2 Lingkaran
	var p Titik

	fmt.Scan(&c1.center.x, &c1.center.y, &c1.radius)

	fmt.Scan(&c2.center.x, &c2.center.y, &c2.radius)

	fmt.Scan(&p.x, &p.y)

	dalam1 := didalam(c1, p)
	dalam2 := didalam(c2, p)

	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
