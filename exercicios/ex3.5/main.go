package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
)

//implemente o conjunto de mendelbrot todo colorido
// usando a função NewRGBA e o tipo color.RGBA ou color.YCbCr.

const (
	mandelWidth, mandelHeight = 800, 800

	xmin, ymin = -2.0, -1.5
	xmax, ymax = 1.0, 1.5

	iterations = 200

	escapeRadiusSq = 4.0
)

func main() {
	http.HandleFunc("/", mandelbrotHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func mandelbrotHandler(w http.ResponseWriter, h *http.Request) {
	w.Header().Set("Content-Type", "image/png")

	img := image.NewRGBA(image.Rect(0, 0, mandelWidth, mandelHeight))

	for py := range mandelHeight {
		var y float64 = float64(py)/mandelHeight*(ymax-ymin) + ymin
		for px := range mandelWidth {
			var x float64 = float64(px)/mandelWidth*(xmax-xmin) + xmin

			c := complex(x, y)

			img.Set(px, py, mandelbrotColor(c))
		}
	}
	png.Encode(w, img)
}

func mandelbrotColor(c complex128) color.Color {
	z := 0 + 0i

	for n := range iterations {
		z = z*z + c

		if cmplx.Abs(z*z) > escapeRadiusSq {
			return colorFromIterations(n)
		}
	}

	return color.RGBA{A: 255}
}

func colorFromIterations(n int) color.Color {
	r := uint8(n % 16 * 16)
	g := uint8(n % 8 * 32)
	b := uint8(n % 4 * 64)

	return color.RGBA{r, g, b, 255}
}
