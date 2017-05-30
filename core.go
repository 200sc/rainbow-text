package main

import (
	"fmt"
	"image/color"

	"github.com/200sc/go-dist/floatrange"

	"image"

	"bitbucket.org/oakmoundstudio/oak"
	"bitbucket.org/oakmoundstudio/oak/render"
)

var (
	font       *render.Font
	r, g, b, a float64
	diff       = floatrange.NewSpread(0, 10)
	limit      = floatrange.NewLinear(0, 255)
)

func main() {
	oak.AddScene("demo",
		// Init
		func(prevScene string, payload interface{}) {
			fg := render.FontGenerator{
				File:    "luxisr.ttf",
				Color:   image.NewUniform(color.RGBA{255, 0, 0, 255}),
				Size:    400,
				Hinting: "",
				DPI:     10,
			}
			r = 255
			font = fg.Generate()
			txt := font.NewStrText("Rainbow", 200, 200)
			render.Draw(txt, 0)
		},
		// Loop
		func() bool {
			r = limit.EnforceRange(r + diff.Poll())
			g = limit.EnforceRange(g + diff.Poll())
			b = limit.EnforceRange(b + diff.Poll())
			fmt.Println(uint8(r), uint8(g), uint8(b))
			// This should be a function in oak to just set color source
			// (or texture source)
			font.Drawer.Src = image.NewUniform(
				color.RGBA{
					uint8(r),
					uint8(b),
					uint8(g),
					255,
				},
			)
			return true
		},

		// End
		func() (string, *oak.SceneResult) {
			return "demo", nil
		},
	)
	render.SetDrawStack(
		render.NewHeap(false),
		render.NewDrawFPS(),
	)
	oak.Init("demo")
}
