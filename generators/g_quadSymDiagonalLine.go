package generators

import (
	"fmt"
	"image"
	"image/color"
	"os"

	"../tools"
)

var quadSymDiagonalLineAdditionalDraw func(*image.RGBA64, int, int, int, int, int, *color.RGBA64)

func quadSymDiagonalLine(coef int, colors []color.RGBA64) {
	if cycles < 0 {
		cycles = 10000000
	}

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

		x = prevX
		y = prevY

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
			Line(img, prevX, bounds.Max.Y-prevY, x, bounds.Max.Y-y, currentColor)
			Line(img, bounds.Max.X-prevX, prevY, bounds.Max.X-x, y, currentColor)
			Line(img, bounds.Max.X-prevX, bounds.Max.Y-prevY, bounds.Max.X-x, bounds.Max.Y-y, currentColor)

			if quadSymDiagonalLineAdditionalDraw != nil {
				quadSymDiagonalLineAdditionalDraw(img, prevX, prevY, x, y, coef, &currentColor)
				quadSymDiagonalLineAdditionalDraw(img, prevX, bounds.Max.Y-prevY, x, bounds.Max.Y-y, coef, &currentColor)
				quadSymDiagonalLineAdditionalDraw(img, bounds.Max.X-prevX, prevY, bounds.Max.X-x, y, coef, &currentColor)
				quadSymDiagonalLineAdditionalDraw(img, bounds.Max.X-prevX, bounds.Max.Y-prevY, bounds.Max.X-x, bounds.Max.Y-y, coef, &currentColor)
			}
		}

		if updateImage != nil && needUpdate {
			updateImage(img)
		}

		prevX = x
		prevY = y
	}
}

func extendableQuadSymDiagonalLine(coef int, colors []color.RGBA64,
	additionalDraw func(*image.RGBA64, int, int, int, int, int, *color.RGBA64)) {
	quadSymDiagonalLineAdditionalDraw = additionalDraw
	quadSymDiagonalLine(coef, colors)
}

// QuadSymDiagonalLine draws an image with diagonal lines with quad symetry
func QuadSymDiagonalLine() {
	if param1 == -1 {
		fmt.Println("p1 must be set for BDL4")
		os.Exit(-1)
	}

	if colors == nil {
		fmt.Println("colors must be set for BDL4")
		os.Exit(-1)
	}

	quadSymDiagonalLineAdditionalDraw = nil
	quadSymDiagonalLine(param1, colors)
}
