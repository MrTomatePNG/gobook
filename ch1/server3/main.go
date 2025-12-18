package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler responde às requisições HTTP gerando figuras de Lissajous
func handler(w http.ResponseWriter, r *http.Request) {
	// Obtém os parâmetros da URL com valores padrão
	cycles := getIntParam(r, "cycles", 5)
	res := getFloatParam(r, "res", 0.001)
	size := getIntParam(r, "size", 100)
	nframes := getIntParam(r, "nframes", 64)
	delay := getIntParam(r, "delay", 8)

	// Gera a figura de Lissajous com os parâmetros especificados
	lissajous(w, cycles, res, size, nframes, delay)
}

// getIntParam extrai um parâmetro inteiro da URL ou retorna o valor padrão
func getIntParam(r *http.Request, param string, defaultValue int) int {
	valueStr := r.URL.Query().Get(param)
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Printf("Parâmetro inválido %s=%s, usando padrão %d\n", param, valueStr, defaultValue)
		return defaultValue
	}

	return value
}

// getFloatParam extrai um parâmetro float64 da URL ou retorna o valor padrão
func getFloatParam(r *http.Request, param string, defaultValue float64) float64 {
	valueStr := r.URL.Query().Get(param)
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		log.Printf("Parâmetro inválido %s=%s, usando padrão %f\n", param, valueStr, defaultValue)
		return defaultValue
	}

	return value
}

// lissajous gera uma animação GIF de curvas de Lissajous aleatórias
func lissajous(out io.Writer, cycles int, res float64, size, nframes, delay int) {
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
