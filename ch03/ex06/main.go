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
		epsX	= (xmax - xmin) / width
		epsY	= (ymax - ymin) / width
	)

	offX := []float64{-epsX, epsX}
	offY := []float64{-epsY, epsY}
 
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin

			subPixels := make([]color.Color, 0)
			for i :=0; i < 2; i++ {
				for j := 0; j < 2; j++ {
					//z := complex(x, y)
					z := complex(x+offX[i], y+offY[i])
					subPixels = append(subPixels, mandelbrot(z))
				}
			}

			img.Set(px, py, avgColor(subPixels))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func avgColor(colors []color.Color) color.Color {
	var r, g, b, a uint16
	n := len(colors)
	for _, c := range colors {
		r_, g_, b_, a_ := c.RGBA()
		r += uint16(r_ / uint32(n))
		g += uint16(g_ / uint32(n))
		b += uint16(b_ / uint32(n))
		a += uint16(a_ / uint32(n))
	}
	return color.RGBA64{r, g, b, a}
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
