package generators

import (
	"fmt"
	"image/color"

	"../tools"
)

func computeLines(nbLines int) [][]int {
	lines := [][]int{}
	generatorx := tools.NewNumberGenerator(seed, 0, img.Bounds().Max.X)
	generatory := tools.NewNumberGenerator(seed, 0, img.Bounds().Max.Y)

	for i := 0; i < nbLines; i++ {
		lines = append(
			lines,
			[]int{
				generatorx.NextPositive(),
				generatory.NextPositive(),
				generatorx.NextPositive(),
				generatory.NextPositive(),
			},
		)
	}

	return lines
}

// AutomoveLines draws a diamond grid
func AutomoveLines(coef int, nbLines int, c color.RGBA64) {
	fmt.Println(c, seed)
	generator := tools.NewNumberGenerator(seed, 0, 8)

	lines := computeLines(nbLines)
	bounds := img.Bounds()

	xOk, yOk, needUpdate := false, false, false

	for i := 0; i < 5000; i++ {
		for j := 0; j < nbLines; j++ {
			xOk = false
			yOk = false
			needUpdate = false

			basicMoveLine(&lines[j][0], &lines[j][1], &lines[j][2], &lines[j][3], coef, &generator)

			if lines[j][2] > bounds.Max.X {
				lines[j][2] = bounds.Max.X
			} else if lines[j][2] < 0 {
				lines[j][2] = 0
			} else {
				xOk = true
			}

			if lines[j][3] > bounds.Max.Y {
				lines[j][3] = bounds.Max.Y
			} else if lines[j][3] < 0 {
				lines[j][3] = 0
			} else {
				yOk = true
			}

			if xOk && yOk {
				needUpdate = true

				AddToLine(img, lines[j][0], lines[j][1], lines[j][2], lines[j][3], c)
			}
		}

		if updateImage != nil && needUpdate {
			updateImage(img)
		}
	}
}
