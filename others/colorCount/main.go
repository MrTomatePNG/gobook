package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"
	"sort"
)

// Quantiza cores reduzindo a precisão (agrupa cores similares)
func quantizeColor(r, g, b uint8, levels int) uint32 {
	divisor := uint8(256 / levels)
	rq := (r / divisor) * divisor
	gq := (g / divisor) * divisor
	bq := (b / divisor) * divisor
	return uint32(rq)<<16 | uint32(gq)<<8 | uint32(bq)
}

// Calcula distância euclidiana entre duas cores
func colorDistance(r1, g1, b1, r2, g2, b2 uint8) float64 {
	dr := float64(r1) - float64(r2)
	dg := float64(g1) - float64(g2)
	db := float64(b1) - float64(b2)
	return math.Sqrt(dr*dr + dg*dg + db*db)
}

type Color struct {
	R, G, B uint8
	Count   int
}

func getDominantColors(imgPath string, numColors, quantLevels int) ([]Color, error) {
	file, err := os.Open(imgPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	colorMap := make(map[uint32]int)
	bounds := img.Bounds()

	// Conta cores quantizadas
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			r8, g8, b8 := uint8(r>>8), uint8(g>>8), uint8(b>>8)

			key := quantizeColor(r8, g8, b8, quantLevels)
			colorMap[key]++
		}
	}

	// Converte para slice e ordena
	colors := make([]Color, 0, len(colorMap))
	for key, count := range colorMap {
		colors = append(colors, Color{
			R:     uint8(key >> 16),
			G:     uint8((key >> 8) & 0xFF),
			B:     uint8(key & 0xFF),
			Count: count,
		})
	}

	sort.Slice(colors, func(i, j int) bool {
		return colors[i].Count > colors[j].Count
	})

	// Retorna as N cores mais dominantes
	if len(colors) > numColors {
		colors = colors[:numColors]
	}

	return colors, nil
}

func main() {
	// quantLevels: quanto menor, mais cores similares são agrupadas
	// 8 níveis = agrupar em 32 tons por canal
	colors, err := getDominantColors("/home/tomate/Imagens/wallpapers/diana-pragmata-hugh-3840x2160-24938.jpg", 5, 8)
	if err != nil {
		panic(err)
	}

	fmt.Println("Cores dominantes:")
	for i, c := range colors {
		fmt.Printf("%d. RGB(%3d, %3d, %3d) - %d pixels - #%02X%02X%02X\n",
			i+1, c.R, c.G, c.B, c.Count, c.R, c.G, c.B)
	}
}
