package generators

import (
	"fmt"
	"image"
	"image/color"
	"os"

	"kirina/tools"
)

func inversedSurroundedDiagonalLineAdditionalDraw(img *image.RGBA64, prevX, prevY, x, y, coef int, color *color.RGBA64) {
	maxSurround := coef / 2
	currentColor := *color

	Line(img, prevX, prevY, x, y, baseColor)

	if (x > prevX && y > prevY) || (x < prevX && y < prevY) {
		for diff := 1; diff < maxSurround; diff++ {
			currentColor = tools.ColorMean(
				100*(diff)/maxSurround, color, &baseColor)

			Line(img, prevX, prevY+diff, x-diff, y, currentColor)
			Line(img, prevX+diff, prevY, x, y-diff, currentColor)
		}
	} else {
		for diff := 1; diff < maxSurround; diff++ {
			currentColor = tools.ColorMean(
				100*(diff)/(maxSurround), color, &baseColor)

			Line(img, prevX, prevY-diff, x-diff, y, currentColor)
			Line(img, prevX+diff, prevY, x, y+diff, currentColor)
		}
	}
}

// InversedSurroundedDiagonalLine draws an image with diagonal lines
func InversedSurroundedDiagonalLine() {
	if param1 == -1 {
		fmt.Println("p1 must be set for IBDL")
		os.Exit(-1)
	}

	if colors == nil {
		fmt.Println("colors must be set for IBDL")
		os.Exit(-1)
	}

	extendableDiagonalLine(param1, colors, inversedSurroundedDiagonalLineAdditionalDraw)
}

// BiSymInversedSurroundedDiagonalLine draws an image with diagonal lines and bi symetry
func BiSymInversedSurroundedDiagonalLine() {
	if param1 == -1 {
		fmt.Println("p1 must be set for IBDL2")
		os.Exit(-1)
	}

	if colors == nil {
		fmt.Println("colors must be set for IBDL2")
		os.Exit(-1)
	}

	if param2 != 0 && param2 != 1 && param2 != 2 {
		fmt.Println("p2 must be set for IBDL2 with 0 (horizontal), 1 (vertical) or 2 (diagonal)")
		os.Exit(-1)
	}

	extendableBiSymDiagonalLine(param1, colors, param2, inversedSurroundedDiagonalLineAdditionalDraw)
}

// QuadSymInversedSurroundedDiagonalLine draws an image with diagonal lines and quad symetry
func QuadSymInversedSurroundedDiagonalLine() {
	if param1 == -1 {
		fmt.Println("p1 must be set for IBDL4")
		os.Exit(-1)
	}

	if colors == nil {
		fmt.Println("colors must be set for IBDL4")
		os.Exit(-1)
	}

	extendableQuadSymDiagonalLine(param1, colors, inversedSurroundedDiagonalLineAdditionalDraw)
}
