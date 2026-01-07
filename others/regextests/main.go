package main

import (
	"fmt"
	"regexp/syntax"
)

func main() {
	pattern := `a+b`

	re, _ := syntax.Parse(pattern, syntax.Perl)
	prog, _ := syntax.Compile(re)

	fmt.Println(prog)
}
