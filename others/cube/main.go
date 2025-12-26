package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	"golang.org/x/term"
)

func main() {
	var A, B, C float64 //angulos de rotação para os eixos X Y e Z

	width, height, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Erro ao obter dimensões:", err)
		return
	}

	fmt.Printf("Largura: %d, Altura: %d\n", width, height)
	zBuffer := make([]float64, width*height)
	buffer := make([]string, width*height)

	fmt.Print("\033[2j") // limpa a tela

	for {

		for i := range buffer {
			buffer[i] = " "
			zBuffer[i] = 0
		}

		for cubeX := -(10.0 * 2); cubeX < (10. * 2); cubeX += 0.5 {
			for cubeY := -(10.0 * 2); cubeY < (10.00 * 2); cubeY += 0.5 {
				calculateForSurface(cubeX, cubeY, (10 * 2), "@", &A, &B, &C, width, height, buffer, zBuffer)
				calculateForSurface(cubeX, cubeY, -(10 * 2), "#", &A, &B, &C, width, height, buffer, zBuffer)
				calculateForSurface((10 * 2), cubeY, cubeX, "$", &A, &B, &C, width, height, buffer, zBuffer)
				calculateForSurface(-(10 * 2), cubeY, cubeX, "~", &A, &B, &C, width, height, buffer, zBuffer)
			}
		}

		fmt.Print("\033[H")
		fmt.Println(strings.Join(buffer, ""))

		A += 0.05 // Incrementa rotação
		B += 0.05
		time.Sleep(16 * time.Millisecond)
		width, height, _ = term.GetSize(int(os.Stdin.Fd()))
		zBuffer = make([]float64, width*height)
		buffer = make([]string, width*height)
	}

}

func calculateForSurface(x, y, z float64, char string, A, B, C *float64, width, height int, buffer []string, zBuffer []float64) {
	radA, radB, radC := *A, *B, *C

	newX := y*math.Sin(radA)*math.Sin(radB)*math.Cos(radC) -
		x*math.Cos(radA)*math.Sin(radB)*math.Cos(radC) + y*math.Cos(radA)*math.Sin(radC) + z*math.Sin(radA)*math.Sin(radC) +
		x*math.Cos(radB)*math.Cos(radC)

	newY := y*math.Cos(radA)*math.Cos(radC) + z*math.Sin(radA)*math.Cos(radC) -
		y*math.Sin(radA)*math.Sin(radB)*math.Sin(radC) + z*math.Cos(radA)*math.Sin(radB)*math.Sin(radC) -
		x*math.Cos(radB)*math.Sin(radC)

	newZ := z*math.Cos(radA)*math.Cos(radB) - y*math.Sin(radA)*math.Cos(radB) + x*math.Sin(radB)

	K1 := 40.
	ooz := 1 / (newZ + 100)

	xp := int(float64(width/2) + K1*newX*ooz*2)
	yp := int(float64(height/2) - K1*newY*ooz)

	if xp >= 0 && xp < width && yp >= 0 && yp < height {
		idx := xp + yp*width
		if ooz > zBuffer[idx] {
			zBuffer[idx] = ooz
			buffer[idx] = char
		}
	}
}
