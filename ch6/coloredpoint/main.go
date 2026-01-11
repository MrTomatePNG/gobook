package main

import (
	"fmt"
	"gobook/ch6/geometry"
	"image/color"
	"sync"
)

type ColoredPoint struct {
	geometry.Point
	Color color.RGBA
}

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func main() {
	var cp ColoredPoint
	metodo := cp.Distance
	fmt.Println(metodo(cp.Point))
	cp.X = 1
	fmt.Println(cp.Point.X)
	cp.Point.Y = 2
	fmt.Println(cp.Y)

	anotherFUnc(metodo)
}

func anotherFUnc(f func(q geometry.Point) float64) {
	f(geometry.Point{1, 2})
}
