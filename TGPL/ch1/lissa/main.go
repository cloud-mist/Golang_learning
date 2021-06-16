package main

// TODO: 练习5,6没写  <16-06-21, cloud mist> //
import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

// 复合声明,生成一个slice切片
var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissa(os.Stdout)
}

func lissa(out io.Writer) {
	// 声明常量
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 9
	)

	freq := rand.Float64() * 3.0
	// 复合声明，gif.GIF 是一个struct类型
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		// 循环64次，每次生成一个201*201的图片，白色和黑色。
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)

	}
	gif.EncodeAll(out, &anim)
}
