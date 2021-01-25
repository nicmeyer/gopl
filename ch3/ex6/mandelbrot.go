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

	var imgL [2 * width][2 * height]color.Color
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
			a1, a2, a3, a4 := imgL[px][py].RGBA()
			b1, b2, b3, b4 := imgL[px+1][py].RGBA()
			c1, c2, c3, c4 := imgL[px][py+1].RGBA()
			d1, d2, d3, d4 := imgL[px+1][py+1].RGBA()

			R := int(a1) + int(b1) + int(c1) + int(d1)
			G := int(a2) + int(b2) + int(c2) + int(d2)
			B := int(a3) + int(b3) + int(c3) + int(d3)
			A := int(a4) + int(b4) + int(c4) + int(d4)

			imgH.Set(px/2, py/2, color.RGBA{
				uint8(R / 1028), uint8(G / 1028), uint8(B / 1028), uint8(A / 1028),
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
