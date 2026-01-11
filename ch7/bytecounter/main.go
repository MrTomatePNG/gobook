package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (int, error) {
	*b += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("Hello"))
	fmt.Println(c)
	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)

	fmt.Println(c)

	var w io.Writer

	w = os.Stdout
	w = new(bytes.Buffer)
	fmt.Println(w)
}
