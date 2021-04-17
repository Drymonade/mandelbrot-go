package main

import (
	"image"
	"image/color/palette"
	"image/gif"
	"math/cmplx"
	"os"
)

func main() {
	const (
		width, height = 1024, 1024
		nframes       = 200
		zoom          = 0.98
	)

	var xmin, ymin, xmax, ymax = -2.0, -2.0, 2.0, 2.0

	anim := gif.GIF{}

	for n := 0; n < nframes; n++ {
		img := image.NewPaletted(
			image.Rect(0, 0, width, height),
			palette.WebSafe,
		)

		for py := 0; py < height; py++ {
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				img.Set(px, py, palette.WebSafe[mandelbrot(z)])
			}
		}

		anim.Delay = append(anim.Delay, 3)
		anim.Image = append(anim.Image, img)

		xmin = (xmin+1.5)*zoom - 1.5
		xmax = (xmax+1.5)*zoom - 1.5

		ymin *= zoom
		ymax *= zoom

	}

	gif.EncodeAll(os.Stdout, &anim)

}

func mandelbrot(z complex128) uint8 {
	const iterations = 215

	var v complex128

	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return n
		}
	}
	return 0
}
