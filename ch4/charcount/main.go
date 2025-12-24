package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

// As vezes precisamos de um mapa ou de um conjunto cujas chaves sejam fatias,
// mas pelo fato de as chaves de uma mapa terem de ser comparaveis, isso não
// pode ser expresso diretamente.
// No entanto, pode ser feito em dois passos. Inicialmente definimos uma chave,
// auxiliar "k" que mapeia cada chave a uma string, coma a propriedade k(x) == k(y)
// se e somente se consederarmos x e y equivalentes. Em seguida, criamos um mapa cujas chaves, aplicando a função auxiliar a cada chave antes de acessarmos o mapa.

var m = make(map[string]int)

func k(list []string) string { return fmt.Sprintf("%q", list) }

func Add(list []string) { m[k(list)]++ }

func Count(list []string) int { return m[k(list)] }

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int

	invalid := 0

	in := bufio.NewReader(os.Stdin)

	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "CHARCOUNT: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Printf("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
