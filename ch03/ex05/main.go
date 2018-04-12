package main
 
import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)
 
func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height	  = 1024, 1024
	)
 
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}
 
func mandelbrot(z complex128) color.Color {
	const contrast = 17
	const iterations = 255 / contrast
	var v complex128

    var nred, ngreen, nblue uint8
 
    for n := uint8(0); n <= iterations; n++ {
		nc := n * contrast

		if(nc < 127) {
			nred = 0
			nblue = 255 - (nc * 2)
			ngreen = nc * 2

		} else { // nc >= 128
			nred = (nc - 128) * 2 
			nblue = 0
			ngreen = 255 - (nc - 128) * 2
 
			v = v*v + z
			if cmplx.Abs(v) > 3 {
				return color.RGBA{nred, ngreen, nblue, 255}
			}
		}
	}
	return color.Black
}
