// Web server that writes the lissajous image to the HTTP client
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette2 = []color.Color{color.White, color.Black, color.RGBA{0x00, 0xff, 0x00, 0x00}, color.RGBA{0xff, 0xff, 0x00, 0x00}, color.RGBA{0x033, 0x54, 0x64, 0x00}}

const (
	whiteIndex2  = 0 // first color in palette
	blackIndex2  = 1 // next color in palette
	greenIndex2  = 2
	color2Index2 = 3
	color3Index2 = 4
)

func main() {
	//handler := func(w http.ResponseWriter, r *http.Request) {
	//	lissajous(w)
	//}
	//http.HandleFunc("/", handler)
	//or
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cycles, err := strconv.Atoi(r.URL.Query().Get("cycles"))
		if err != nil {
			w.Write([]byte("Error with cycles parameter"))
			return
		}
		lissajous2(w, cycles)
	})
	http.ListenAndServe("localhost:8000", nil)
}

func lissajous2(out io.Writer, cycles int) {
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
		img := image.NewPaletted(rect, palette2)

		cindex := uint8(0)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			if cindex > 4 {
				cindex = 0
			}
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), cindex)
			cindex++
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
