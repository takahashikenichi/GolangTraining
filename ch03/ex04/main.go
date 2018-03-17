// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"errors"
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
	colorTop	  = "#ff0000"
	color_1	  = "#dd0022"
	color_2	  = "#bb0055"
	color_3	  = "#990077"
	color_4	  = "#770099"
	color_5	  = "#5500bb"
	color_6	  = "#2200dd"
	colorBottom   = "#0000ff"
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, z, error := corner(i+1, j)
			if error != nil { continue }
			bx, by, z, error := corner(i, j)
			if error != nil { continue }
			cx, cy, z, error := corner(i, j+1)
			if error != nil { continue }
			dx, dy, z, error := corner(i+1, j+1)
			if error != nil { continue }

			color := "#000000"
			switch {
			case z > 0.6:
				color = colorTop
			case z > 0.3 && z <= 0.6:
				color = color_1
			case z > 0.0 && z <= 0.3:
				color = color_2
			case z > -0.3 && z <= 0.0:
				color = color_3
			case z > -0.3 && z <= 0.0:
				color = color_4
			case z > -0.6 && z <= -0.3:
				color = color_5
			case z <= -0.6:
				color = color_6
			default:
				color = colorBottom
			}
			
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64, error) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	if math.IsNaN(z) || math.IsInf(z, 0) {
		return 0, 0, 0, errors.New("Invalid value:z")
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, nil
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

