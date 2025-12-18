// Segundo a abordagem do exemplo de Lissoujous na seção 1.7.
// Crie um servidor web que calcule superficies e escreva dados SVG ao cliente
// O servidor deve definir o cabeçalho content-type para o 'image/svg+xml'
// Permita que cliente especifique valores como altura,largura e cor como parametros
// da Requisição HTTP

package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)
const Z_MAX = 1.0
const Z_MIN = -0.5

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	currentWidth := width   // Valor padrão: 600
	currentHeight := height // Valor padrão: 320

	if s := r.URL.Query().Get("width"); s != "" {
		if v, err := strconv.Atoi(s); err == nil && v > 0 {
			currentWidth = v
		}
	}

	if s := r.URL.Query().Get("height"); s != "" {
		if v, err := strconv.Atoi(s); err == nil && v > 0 {
			currentHeight = v
		}
	}
	currentXyScale := float64(currentWidth) / 2 / xyrange
	currentZscale := float64(currentHeight) * 0.4

	t := 0.0

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke:black; fill:none; stroke-width:0.1' "+
		"width='%d' height='%d'>\n", currentWidth, currentHeight)
	for i := range cells {
		for j := range cells {
			ax, ay := corner(i+1, j, t, currentWidth, currentHeight, currentXyScale, currentZscale)
			bx, by := corner(i, j, t, currentWidth, currentHeight, currentXyScale, currentZscale)
			cx, cy := corner(i, j+1, t, currentWidth, currentHeight, currentXyScale, currentZscale)
			dx, dy := corner(i+1, j+1, t, currentWidth, currentHeight, currentXyScale, currentZscale)

			x_for_z := xyrange * (float64(i)/cells - 0.5)
			y_for_z := xyrange * (float64(j)/cells - 0.5)
			z_val := f(x_for_z, y_for_z, t)
			polygonColor := color(z_val)

			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s' stroke='none'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, polygonColor)
		}
	}
	fmt.Fprintf(w, "</svg>\n")
}

func corner(i, j int, t float64, w, h int, xs, zs float64) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y, t)

	sx := float64(w)/2 + (x-y)*cos30*xs
	sy := float64(h)/2 + (x+y)*sin30*xs - z*zs

	return sx, sy
}

func f(x, y, t float64) float64 {
	r := math.Hypot(x, y)
	if r == 0 {
		return 1.0
	}
	return math.Sin(r+t) / r
}

func color(z float64) string {
	if z > Z_MAX {
		z = Z_MAX
	}
	if z < Z_MIN {
		z = Z_MIN
	}

	normZ := (z - Z_MIN) / (Z_MAX - Z_MIN)

	B := int(255 * (1 - normZ))

	R := int(255 * normZ)

	G := 0

	return fmt.Sprintf("#%02X%02X%02X", R, G, B)
}
