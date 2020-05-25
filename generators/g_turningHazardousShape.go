package generators

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"os"

	"../tools"
)

func drawTurningSandCurve(img *image.RGBA64, points []Point, c color.RGBA64, sandCoef int) {
	for index, point := range points {
		if index+1 == len(points) {
			addToTurningSandLine(img, point, points[0], c, sandCoef)
		} else {
			addToTurningSandLine(img, point, points[index+1], c, sandCoef)
		}
	}
}

// Custom version of AddToSandLine with a strange point computation
func addToTurningSandLine(img *image.RGBA64, p0, p1 Point, c color.RGBA64, sandCoef int) {
	pt := PolarPoint{
		angle:  p0.polar.angle + p1.polar.angle/2,
		radius: p0.polar.radius + float64(sandCoef),
	}

	pr := pt.toCartesian(p0.center)

	// implemented straight from WP pseudocode
	dx := p1.cartesian.x - p0.cartesian.x
	if dx < 0 {
		dx = -dx
	}

	dy := p1.cartesian.y - p0.cartesian.y
	if dy < 0 {
		dy = -dy
	}

	var sx, sy int
	if p0.cartesian.x < p1.cartesian.x {
		sx = 1
	} else {
		sx = -1
	}

	if p0.cartesian.y < p1.cartesian.y {
		sy = 1
	} else {
		sy = -1
	}

	err := dx - dy

	for {
		nextXMove := 0
		nextYMove := 0

		e2 := 2 * err

		if e2 > -dy {
			err -= dy
			nextXMove = sx
		}

		if e2 < dx {
			err += dx
			nextYMove = sy
		}

		p0.cartesianMove(nextXMove, nextYMove)

		if param2 == 1 {
			pr.x += nextXMove
			pr.y += nextYMove
		}

		AddToDegressiveLine(img, p0.cartesian.x, p0.cartesian.y, pr.x, pr.y, c)

		if p0.cartesian.x == p1.cartesian.x && p0.cartesian.y == p1.cartesian.y {
			break
		}
	}
}

func thsCreateColor(i int) color.RGBA64 {
	nbColors := cycles
	nbAvailableColors := len(colors)

	nbColorsByIndex := nbColors / nbAvailableColors
	currentIndex := i * nbAvailableColors / nbColors
	nextIndex := currentIndex + 1

	if currentIndex == nbAvailableColors-1 {
		nextIndex = 0
	}

	value := (i - (nbColorsByIndex * currentIndex)) * 100 / nbColorsByIndex

	computedColor := tools.ColorMean(value, &colors[currentIndex], &colors[nextIndex])

	computedColor.R /= uint16(cycles)
	computedColor.G /= uint16(cycles)
	computedColor.B /= uint16(cycles)

	return computedColor
}

// TurningHazardousShape draws a turning shape
func TurningHazardousShape() {
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
	generatorPoints := tools.NewNumberGenerator(seed, 40, 70)
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

		color := thsCreateColor(i)

		drawTurningSandCurve(img, shape, color, param1)

		if updateImage != nil {
			updateImage(img)
		}
	}
}
