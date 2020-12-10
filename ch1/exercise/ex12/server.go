// Server1 is a minimal "echo" server.
// Lissajous generates GIF animations of random Lissajous figures.
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
)

var palette = []color.Color{color.Black, color.White, color.RGBA{0x00, 0xff, 0x00, 0xff}}

const (
	blackIndex = 0 // first color in palette
	whiteIndex = 1 // next color in palette
	greenIndex = 2 // last color in palette
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		var cycles int // number of complete x oscillator revolutions
		cyclesStr := r.FormValue("cycles")
		if cyclesStr == "" {
			cycles = 5
		} else {
			var err error
			cycles, err = strconv.Atoi(cyclesStr)
			if err != nil {
				log.Print(err)
			}
		}
		lissajous(w, cycles)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, cycles int) {
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			if x > y {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
					whiteIndex)
			} else {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
					greenIndex)
			}
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
