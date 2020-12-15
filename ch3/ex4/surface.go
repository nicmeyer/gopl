// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	cells   = 100         // number of grid cells
	xyrange = 30.0        // axis ranges (-xyrange..+xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)
var width, height = 600.0, 320.0                    // canvas size in pixels
var xyscale = width / 2 / xyrange                   // pixels per x or y unit
var zscale = height * 0.4                           // pixels per z unit
var colors = [2]string{"#0000ff", "#ff0000"}

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		for k, v := range r.Form {
			switch k {
			case "width":
				width = parseFloat(v[0])
			case "height":
				height = parseFloat(v[0])
			case "color":
				for i, c := range strings.Split(v[0], ",") {
					colors[i] = c
				}
			}
		}
		w.Header().Set("Content-Type", "image/svg+xml")
		fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: grey; fill: white; stroke-width: 0.7' "+
			"width='%f' height='%f'>", width, height)
		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				ax, ay := corner(i+1, j)
				bx, by := corner(i, j)
				cx, cy := corner(i, j+1)
				dx, dy := corner(i+1, j+1)

				x, y := translate(i, j)
				z := f(x, y)
				color := colors[0]
				if z >= 0 {
					color = colors[1]
				}

				fmt.Fprintf(w, "<polygon style='fill: %s' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					color, ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
		fmt.Fprintln(w, "</svg>")
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func parseFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "surface: %v\n", err)
		os.Exit(1)
	}
	return f
}

func translate(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	return x, y
}

func corner(i, j int) (float64, float64) {
	x, y := translate(i, j)
	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
