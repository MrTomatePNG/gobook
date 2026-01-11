package main

import (
	"flag"
	"fmt"
	"gobook/ch7/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.00, "A temperatura")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
