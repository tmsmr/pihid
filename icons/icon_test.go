package icons

import (
	"testing"
)

import "github.com/fogleman/gg"

func TestDrawIcon(t *testing.T) {
	dc := gg.NewContext(128, 128)
	dx := 32
	dy := 32
	for y := 0; y < 64; y++ {
		for x := 0; x < 64; x++ {
			a := Power.Alpha64x64[y*64+x]
			if a > 0 {
				dc.SetRGBA(1, 1, 1, float64(a)/255)
				dc.SetPixel(x+dx, y+dy)
			}
		}
	}
	err := dc.SavePNG("power_test.png")
	if err != nil {
		panic(err)
	}
}
