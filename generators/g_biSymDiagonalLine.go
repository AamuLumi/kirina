package generators

import (
	"fmt"
	"image"
	"image/color"
	"os"

	"../tools"
)

var biSymDiagonalLineAdditionalDraw func(*image.RGBA64, int, int, int, int, int, *color.RGBA64)

func biSymDiagonalLine(coef int, colors []color.RGBA64, symetryType int) {
	if cycles < 0 {
		cycles = 1000000
	}

	generator := tools.NewNumberGenerator(seed, 0, 4)
	colorGenerator := tools.NewColorGenerator(42, colors)
	bounds := img.Bounds()

	x, y, prevX, prevY := 0, 0, 0, 0
	x2, y2, prevX2, prevY2 := 0, 0, 0, 0
	xOk, yOk, needUpdate := false, false, false
	var currentColor color.RGBA64

	var drawSymetry func(*image.RGBA64, int, int, int, int) (int, int, int, int)

	switch symetryType {
	case BiSymetry["horizontal"]:
		drawSymetry = horizontalSym
	case BiSymetry["vertical"]:
		drawSymetry = verticalSym
	case BiSymetry["diagonal"]:
		drawSymetry = diagonalSym
	}

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

			prevX2, prevY2, x2, y2 = drawSymetry(img, prevX, prevY, x, y)

			Line(img, prevX2, prevY2, x2, y2, currentColor)

			if biSymDiagonalLineAdditionalDraw != nil {
				biSymDiagonalLineAdditionalDraw(img, prevX, prevY, x, y, coef, &currentColor)
				biSymDiagonalLineAdditionalDraw(img, prevX2, prevY2, x2, y2, coef, &currentColor)
			}
		}

		if updateImage != nil && needUpdate {
			updateImage(img)
		}

		prevX = x
		prevY = y
	}
}

func extendableBiSymDiagonalLine(coef int, colors []color.RGBA64, symetryType int,
	additionalDraw func(*image.RGBA64, int, int, int, int, int, *color.RGBA64)) {
	biSymDiagonalLineAdditionalDraw = additionalDraw
	biSymDiagonalLine(coef, colors, symetryType)
}

// BiSymDiagonalLine draws an image with diagonal lines, with a simple symetry
func BiSymDiagonalLine() {
	if param1 == -1 {
		fmt.Println("p1 must be set for BDL2")
		os.Exit(-1)
	}

	if colors == nil {
		fmt.Println("colors must be set for BDL2")
		os.Exit(-1)
	}

	if param2 != 0 && param2 != 1 && param2 != 2 {
		fmt.Println("p2 must be set for BDL2 with 0 (horizontal), 1 (vertical) or 2 (diagonal)")
		os.Exit(-1)
	}

	biSymDiagonalLineAdditionalDraw = nil
	biSymDiagonalLine(param1, colors, param2)
}

func horizontalSym(img *image.RGBA64, prevX, prevY, x, y int) (int, int, int, int) {
	bounds := img.Bounds()

	return prevX, bounds.Max.Y - prevY, x, bounds.Max.Y - y
}

func verticalSym(img *image.RGBA64, prevX, prevY, x, y int) (int, int, int, int) {
	bounds := img.Bounds()

	return bounds.Max.X - prevX, prevY, bounds.Max.X - x, y
}

func diagonalSym(img *image.RGBA64, prevX, prevY, x, y int) (int, int, int, int) {
	bounds := img.Bounds()

	return bounds.Max.X - prevX, bounds.Max.Y - prevY, bounds.Max.X - x, bounds.Max.Y - y
}
