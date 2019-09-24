package generators

import (
	"image/color"

	"../tools"
)

// Subdivision divide the image
func Subdivision(coef int, colors []color.RGBA64) {
	if cycles < 0 {
		cycles = 10000000
	}

	generator := tools.NewNumberGenerator(seed, 0, 8)
	colorGenerator := tools.NewColorGenerator(42, colors)
	bounds := img.Bounds()

	x, y := 0, 0
	xOk, yOk := false, false

	for i := 0; i < cycles; i++ {
		xOk = false
		yOk = false

		basic8MoveWithCoef(&x, &y, coef, &generator)

		if x > bounds.Max.X {
			x = bounds.Max.X
		} else if x < 0 {
			x = 0
		} else {
			xOk = true
		}

		if y > bounds.Max.Y {
			y = bounds.Max.Y
		} else if y < 0 {
			y = 0
		} else {
			yOk = true
		}

		if xOk && yOk {
			DiamondBigger(img, x, y, coef/2, colorGenerator.NextColor())
		}
	}
}
