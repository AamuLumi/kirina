package generators

import (
	"fmt"
	"image/color"
	"math"
	"os"

	"../tools"
)

// AutomoveCenterLines draws a diamond grid
func AutomoveCenterLines(coef int, nbLines int, c color.RGBA64) {
	if param1 == -1 {
		fmt.Println("p1 must be set for THS")
		os.Exit(-1)
	}

	if colors == nil {
		fmt.Println("colors must be set for THS")
		os.Exit(-1)
	}

	if cycles < 0 {
		cycles = 200
	}

	generator := tools.NewNumberGenerator(seed, 0, 6)
	generatorPoints := tools.NewNumberGenerator(seed, 20, 50)
	generatorPointRandomness := tools.NewNumberGenerator(seed, 0, 100)
	radius := 40.0

	nbPointsToGenerate := generatorPoints.NextPositive()
	bounds := img.Bounds()
	var shape []Point

	for i := 0; i < nbPointsToGenerate; i++ {
		randomness := float64(generatorPointRandomness.NextPositive()-50) / float64(nbPointsToGenerate) / 100
		point := math.Pi*2*float64(i)/float64(nbPointsToGenerate) + randomness

		x := int(math.Cos(point)*radius) + bounds.Max.X/2
		y := int(math.Sin(point)*radius) + bounds.Max.Y/2

		shape = append(shape, Point{
			polar: PolarPoint{
				angle:  point,
				radius: distance(bounds.Max.X/2, bounds.Max.Y/2, x, y),
			},
			cartesian: CartesianPoint{
				x: x,
				y: y,
			},
			center: CartesianPoint{
				x: bounds.Max.X / 2,
				y: bounds.Max.Y / 2,
			},
		})
	}

	for i := 0; i < cycles; i++ {
		for index := range shape {
			moveAwayFromCenterWithCoef(bounds.Max.X/2, bounds.Max.Y/2, &shape[index], 1, &generatorPointRandomness, &generator)
		}

		drawTurningSandCurve(img, shape, c, param1)

		if updateImage != nil {
			updateImage(img)
		}
	}
}
