package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
)

const (
	cells   = 100         // number of grid cells
	xyrange = 30.0        // axis ranges (-xyrange..+xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)
var xyscale, zscale float64
var gw, gh float64

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")

		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(os.Stderr, "form parse err: %v\n", err)
			os.Exit(1)
		}

		width, _ := strconv.ParseInt(r.FormValue("width"), 10, 64)
		height, _ := strconv.ParseInt(r.FormValue("height"), 10, 64)
		gw = float64(width)
		gh = float64(height)

		xyscale = gw / 2 / xyrange
		zscale = gh * 0.4

		fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: grey; fill: white; stroke-width: 0.7' "+
			"width='%d' height='%d'>", width, height)
		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				ax, ay, ok := corner(i+1, j)
				bx, by, ok := corner(i, j)
				cx, cy, ok := corner(i, j+1)
				dx, dy, ok := corner(i+1, j+1)
				if !ok {
					continue
				}
				fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
		fmt.Fprintln(w, "</svg>")
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func corner(i, j int) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z, ok := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := gw/2 + (x-y)*cos30*xyscale
	sy := gh/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, ok
}

func f(x, y float64) (float64, bool) {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r, !math.IsNaN(r)
}
