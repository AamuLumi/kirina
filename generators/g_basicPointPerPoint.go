package generators

import (
	"fmt"
	"os"

	"kirina/tools"
)

// BasicPointPerPoint draws an random image point per point
func BasicPointPerPoint() {
	if param1 == -1 {
		fmt.Println("p1 must be set for BPPP")
		os.Exit(-1)
	}

	if colors == nil {
		fmt.Println("colors must be set for BPPP")
		os.Exit(-1)
	}

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

		basic8MoveWithCoef(&x, &y, param1, &generator)

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
			DrawPoint(img, x, y, colorGenerator.NextColor())
		}
	}
}
