package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash" // Interface comum para todos os hashes
	"io"
	"os"
)

var arch = flag.String("s", "256", "algoritmo de hash: 256, 384 ou 512")

func main() {
	flag.Parse()

	var h hash.Hash // Criamos uma variável que aceita qualquer tipo de hash

	// Escolhemos qual implementação colocar dentro da interface 'h'
	switch *arch {
	case "384":
		h = sha512.New384()
	case "512":
		h = sha512.New()
	case "256":
		h = sha256.New()
	default:
		fmt.Fprintf(os.Stderr, "Algoritmo inválido: %s\n", *arch)
		os.Exit(1)
	}

	// Copiamos os dados da entrada padrão para o hash
	if _, err := io.Copy(h, os.Stdin); err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao processar entrada: %v\n", err)
		os.Exit(1)
	}

	// Agora sim, pegamos o resultado final
	// h.Sum(nil) finaliza o cálculo e retorna o []byte
	resultado := h.Sum(nil)

	fmt.Printf("Algoritmo: SHA%s\n", *arch)
	fmt.Printf("Hash: %x\n", resultado)
	fmt.Printf("Total de bytes: %d\n", len(resultado))
}
