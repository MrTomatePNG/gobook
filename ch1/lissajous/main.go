package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var pallete = []color.Color{
	color.RGBA{0x00, 0x00, 0x00, 0xff}, // Fundo preto
	color.RGBA{0x00, 0xff, 0x00, 0xff}, // Tom de verde (0x00FF00FF)
	color.RGBA{0xff, 0x00, 0x00, 0xff}, // Vermelho
	color.RGBA{0x00, 0x00, 0xff, 0xff}, // Azul
	color.RGBA{0xff, 0xff, 0x00, 0xff}, // Amarelo
	color.RGBA{0xff, 0x00, 0xff, 0xff}, // Magenta
	color.RGBA{0x00, 0xff, 0xff, 0xff}, // Ciano
}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 10
		res     = 0.001
		size    = 200
		nframes = 256
		delay   = 8
	)

	freq := rand.Float64() * 6.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallete)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(math.Floor(rand.Float64()*8)))
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)

	}
	gif.EncodeAll(out, &anim)

}
