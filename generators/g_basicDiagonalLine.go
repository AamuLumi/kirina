package generators

import (
	"fmt"
	"image"
	"image/color"
	"os"

	"kirina/tools"
)

var basicDiagonalLineAdditionalDraw func(*image.RGBA64, int, int, int, int, int, *color.RGBA64)

func diagonalLine(coef int, colors []color.RGBA64) {
	generator := tools.NewNumberGenerator(seed, 0, 4)
	colorGenerator := tools.NewColorGenerator(42, colors)
	bounds := img.Bounds()

	x, y, prevX, prevY := 0, 0, 0, 0
	xOk, yOk, needUpdate := false, false, false
	var currentColor color.RGBA64

	for i := 0; i < cycles; i++ {
		xOk = false
		yOk = false
		needUpdate = false

		basic4MoveWithCoef(&x, &y, coef, &generator)

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
			needUpdate = true

			currentColor = colorGenerator.NextColor()
			Line(img, prevX, prevY, x, y, currentColor)

			if basicDiagonalLineAdditionalDraw != nil {
				basicDiagonalLineAdditionalDraw(img, prevX, prevY, x, y, coef, &currentColor)
			}
		}

		if updateImage != nil && needUpdate {
			updateImage(img)
		}

		prevX = x
		prevY = y
	}
}

func extendableDiagonalLine(coef int, colors []color.RGBA64,
	additionalDraw func(*image.RGBA64, int, int, int, int, int, *color.RGBA64)) {
	basicDiagonalLineAdditionalDraw = additionalDraw
	diagonalLine(coef, colors)
}

// BasicDiagonalLine draws an image with diagonal lines
func BasicDiagonalLine() {
	if param1 == -1 {
		fmt.Println("p1 must be set for BDL2")
		os.Exit(-1)
	}

	if colors == nil {
		fmt.Println("colors must be set for BDL2")
		os.Exit(-1)
	}

	if cycles < 0 {
		cycles = 1000000
	}

	basicDiagonalLineAdditionalDraw = nil
	diagonalLine(param1, colors)
}
