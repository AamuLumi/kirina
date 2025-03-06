package generators

import (
	"fmt"
	"image"
	"image/color"
	"os"

	"kirina/tools"
)

var basicRoundLineAdditionalDraw func(*image.RGBA64, int, int, int, int, int, *color.RGBA64)

func roundLine(coef int, colors []color.RGBA64) {
	generator := tools.NewNumberGenerator(seed, 0, 4)
	colorGenerator := tools.NewColorGenerator(42, colors)
	bounds := img.Bounds()

	x, y, prevX, prevY := 0, 0, 0, 0
	xOk, yOk := false, false
	var currentColor color.RGBA64

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

		if xOk && yOk && img.At(x, y) == baseColor {
			currentColor = colorGenerator.NextColor()
			DrawPoint(img, x, y, currentColor)

			if basicRoundLineAdditionalDraw != nil {
				basicRoundLineAdditionalDraw(img, prevX, prevY, x, y, coef, &currentColor)
			}
		}

		prevX = x
		prevY = y
	}
}

func extendableRoundLine(coef int, colors []color.RGBA64,
	additionalDraw func(*image.RGBA64, int, int, int, int, int, *color.RGBA64)) {
	basicRoundLineAdditionalDraw = additionalDraw
	roundLine(coef, colors)
}

// BasicRoundLine draws an image with a round line
func BasicRoundLine() {
	if param1 == -1 {
		fmt.Println("p1 must be set for BRL")
		os.Exit(-1)
	}

	if colors == nil {
		fmt.Println("colors must be set for BRL")
		os.Exit(-1)
	}

	if cycles < 0 {
		cycles = 10000000
	}

	basicRoundLineAdditionalDraw = nil
	roundLine(param1, colors)
}
