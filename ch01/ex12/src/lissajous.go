package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var palette = []color.Color{
	color.White,
	color.Black,
	color.RGBA{0xff, 0x0,  0x0, 0xff},
	color.RGBA{0x0,  0xff, 0x0, 0xff},
	color.RGBA{0x0,  0x0,  0xff, 0xff},
	color.RGBA{0xff, 0xff, 0x0, 0xff},
	color.RGBA{0x0,  0xff, 0xff, 0xff},
	color.RGBA{0xff, 0x0,  0xff, 0xff}}

const (
	whiteIndex = 0
	blackIndex = 1
	redIndex   = 2
	greenIndex = 3
	blueIndex  = 4
	yellowIndex =5
	cyanIndex  = 6
	magentaIndex = 7
)


func lissajous(out io.Writer, param_cycles int) {
	const (
		//cycles  = 5     // 発振器xが完了する周回の回数
		res     = 0.001 // 回転の分解能
		size    = 100   // 画像キャンバスは[-size..+size]
		nframes = 64    // アニメーションフレーム数
		delay   = 8     // 10ms単位でのフレーム間の遅延
	)
	cycles := float64(param_cycles)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		// set background color
		for h := rect.Min.Y; h < rect.Max.Y; h++ {
			for v := rect.Min.X; v < rect.Max.X; v++ {
				img.SetColorIndex(v, h, blackIndex)
			}
		}

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
//			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), (uint8)(i * 8 / nframes))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // エンコードエラーを無視
}

