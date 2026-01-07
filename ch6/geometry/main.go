package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

func Distance(q, p Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func tests() {

	p := Point{10, 20}
	q := Point{-10, -20}

	//as duas tem o mesmo efeito mas q.Distance é uma
	// funcao membro de Point
	Distance(q, p)
	q.Distance(q)
}

type Path []Point

func (path Path) Distance() (sum float64) {
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return
}

func main() {
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance() / 2)
}
