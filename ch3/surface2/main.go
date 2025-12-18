package main

import (
	"fmt"
	"math"
	"os"
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
	// --- 1. Constantes de Animação ---
	const numFrames = 50   // Total de quadros
	const frameSpeed = 0.2 // Variação de 't' por quadro (velocidade da onda)

	for frame := range numFrames {

		// 2. Cálculo do tempo 't' para o quadro atual
		t := float64(frame) * frameSpeed

		// 3. Gerar o arquivo SVG para este quadro
		if err := generateSVGFrame(frame, t); err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao gerar quadro %d: %v\n", frame, err)
			return
		}
		fmt.Printf("Quadro gerado: frame%03d.svg\n", frame)
	}

	fmt.Println("\nTodos os quadros SVG foram gerados. Próxima etapa: Converter para GIF.")
}

// Nova função que contém a lógica de impressão (a antiga 'main')
func generateSVGFrame(frame int, t float64) error {

	// Abrir o arquivo de saída (ex: frame000.svg, frame001.svg...)
	filename := fmt.Sprintf("frame%03d.svg", frame)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close() // Garantir que o arquivo é fechado

	// Imprimir o cabeçalho SVG no arquivo
	fmt.Fprintf(file, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke:grey; fill:white; stroke-width:0.7' "+
		"width='%d' height='%d'>\n", width, height)

	// Loop principal da grade
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {

			// As chamadas 'corner' AGORA usam o tempo 't'
			ax, ay := corner(i+1, j, t)
			bx, by := corner(i, j, t)
			cx, cy := corner(i, j+1, t)
			dx, dy := corner(i+1, j+1, t)

			x_for_z := xyrange * (float64(i)/cells - 0.5)
			y_for_z := xyrange * (float64(j)/cells - 0.5)
			z_val := f_cone(x_for_z, y_for_z, t) // Chamando a f(x, y, t) para obter o z

			polygonColor := color(z_val) // Obter a cor com base na altura z
			// Imprimir o polígono no arquivo
			fmt.Fprintf(file, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s' stroke='none'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, polygonColor)
		}
	}

	// Fechar a tag SVG
	fmt.Fprintf(file, "</svg>\n")
	return nil
}

func corner(i, j int, t float64) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f_cone(x, y, t)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy
}

func f(x, y, t float64) float64 {
	r := math.Hypot(x, y)
	if r == 0 {
		return 1.0
	}
	return math.Sin(r+t) / r
}
func f_caixa_ovos(x, y, t float64) float64 {
	const amplitude = 0.5
	const frequencia = 4.0
	return amplitude * (math.Sin(x/frequencia+t) + math.Cos(y/frequencia))
}
func f_sela(x, y, t float64) float64 {
	const fator = 0.05
	return fator * (x*x - y)
}

func f_cone(x, y, t float64) float64 {
	const amplitude = 0.5
	r := math.Hypot(x, y)
	return -amplitude * r
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
