// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	var imgL [2*width][2*height]color.Color
	for py := 0; py < height*2; py++ {
		y := float64(py)/(height*2)*(ymax-ymin) + ymin
		for px := 0; px < width*2; px++ {
			x := float64(px)/(width*2)*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			imgL[px][py] = mandelbrot(z)
		}
	}
	imgH := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height*2; py += 2 {
		for px := 0; px < width*2; px += 2 {
			r1, g1, b1, a1 := imgL[px][py].RGBA()
			r2, g2, b2, a2 := imgL[px+1][py].RGBA()
			r3, g3, b3, a3 := imgL[px][py+1].RGBA()
			r4, g4, b4, a4 := imgL[px+1][py+1].RGBA()

			R := r1 + r2 + r3 + r4
			G := g1 + g2 + g3 + g4
			B := b1 + b2 + b3 + b4
			A := a1 + a2 + a3 + a4

			imgH.Set(px/2, py/2, color.RGBA{
				uint8(R/1028), uint8(G/1028), uint8(B/1028), uint8(A/1028),
			})
		}
	}
	f, err := os.Create("mandelbrot.png")
	if err != nil {
		fmt.Fprintf(os.Stdout, "mandelbrot: %v\n", err)
		os.Exit(1)
	}
	if err := png.Encode(f, imgH); err != nil {
		f.Close()
		fmt.Fprintf(os.Stdout, "mandelbrot: %v\n", err)
		os.Exit(1)
	}
	if err := f.Close(); err != nil {
		fmt.Fprintf(os.Stdout, "mandelbrot: %v\n", err)
		os.Exit(1)
	}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{0, 255, 255, 255 - contrast*n}
		}
	}
	return color.RGBA{0, 255, 255, 255}
}
