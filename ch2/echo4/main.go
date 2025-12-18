package echo4

import (
	"flag"
	"fmt"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func gcd(x int, y int) {
	for y != 0 {
		x, y = y, x%y
		fmt.Println("Valor de X: ", x, "\nValor de Y: ", y)
	}
}
