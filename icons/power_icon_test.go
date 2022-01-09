package icons

import (
	"testing"
)

import "github.com/fogleman/gg"

func TestDrawPowerIcon(t *testing.T) {
	dc := gg.NewContext(128, 128)
	dx := 32
	dy := 32
	for x := 0; x < 64; x++ {
		for y := 0; y < 64; y++ {
			a := Power.Alpha64x64[x*64+y]
			if a > 0 {
				dc.SetRGBA(1, 1, 1, float64(a)/255)
				dc.SetPixel(x+dx, y+dy)
			}
		}
	}
	err := dc.SavePNG("power_icon.png")
	if err != nil {
		panic(err)
	}
}
