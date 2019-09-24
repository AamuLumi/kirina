package generators

import (
	"image"
	"image/color"

	"../tools"
)

func additionalDraw(img *image.RGBA64, prevX, prevY, x, y, coef int, color *color.RGBA64) {
	maxSurround := coef / 2
	currentColor := *color

	if (x > prevX && y > prevY) || (x < prevX && y < prevY) {
		for diff := 1; diff < maxSurround; diff++ {
			currentColor = tools.ColorMean(
				100*(maxSurround-diff)/maxSurround, color, &baseColor)

			Line(img, prevX, prevY+diff, x-diff, y, currentColor)
			Line(img, prevX+diff, prevY, x, y-diff, currentColor)
		}
	} else {
		for diff := 1; diff < maxSurround; diff++ {
			currentColor = tools.ColorMean(
				100*(maxSurround-diff)/(maxSurround), color, &baseColor)

			Line(img, prevX, prevY-diff, x-diff, y, currentColor)
			Line(img, prevX+diff, prevY, x, y+diff, currentColor)
		}
	}
}

// SurroundedDiagonalLine draws an image with diagonal lines
func SurroundedDiagonalLine(coef int, colors []color.RGBA64) {
	extendableDiagonalLine(coef, colors, additionalDraw)
}

// BiSymSurroundedDiagonalLine draws an image with diagonal lines and bi symetry
func BiSymSurroundedDiagonalLine(coef int, colors []color.RGBA64, symetryType int) {
	extendableBiSymDiagonalLine(coef, colors, symetryType, additionalDraw)
}

// QuadSymSurroundedDiagonalLine draws an image with diagonal lines and quad symetry
func QuadSymSurroundedDiagonalLine(coef int, colors []color.RGBA64) {
	extendableQuadSymDiagonalLine(coef, colors, additionalDraw)
}
