package generators

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"os"

	"kirina/tools"
)

func tunnelCreateColor(i int) color.RGBA64 {
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

	computedColor.R /= uint16(cycles) / 16
	computedColor.G /= uint16(cycles) / 16
	computedColor.B /= uint16(cycles) / 16

	return computedColor
}

// AddToSandTunnelLine draws line by Bresenham's algorithm.
func AddToSandTunnelLine(img *image.RGBA64, p0, p1 Point, c color.RGBA64, sandCoef int) {
	p0or := Point{
		cartesian: CartesianPoint{
			x: p0.cartesian.x,
			y: p0.cartesian.y,
		},
		polar:  p0.cartesian.toPolar(p0.center),
		center: p0.center,
	}

	p1or := Point{
		cartesian: CartesianPoint{
			x: p1.cartesian.x,
			y: p1.cartesian.y,
		},
		polar:  p1.cartesian.toPolar(p1.center),
		center: p1.center,
	}

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

	p0or.polarMove(0.2, -float64(sandCoef))
	p1or.polarMove(0.2, -float64(sandCoef))

	dx = p0.cartesian.x - p0or.cartesian.x
	if dx < 0 {
		dx = -dx
	}

	dy = p0.cartesian.y - p0or.cartesian.y
	if dy < 0 {
		dy = -dy
	}

	if p0or.cartesian.x < p0.cartesian.x {
		sx = 1
	} else {
		sx = -1
	}

	if p0or.cartesian.y < p0.cartesian.y {
		sy = 1
	} else {
		sy = -1
	}

	err := dx - dy

	for {
		AddToWuLine(img, float64(p0or.cartesian.x), float64(p0or.cartesian.y), float64(p1or.cartesian.x), float64(p1or.cartesian.y), 1.0, c)

		if p0.cartesian.x == p0or.cartesian.x && p0.cartesian.y == p0or.cartesian.y {
			break
		}

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

		p0or.cartesianMove(nextXMove, nextYMove)
		p1or.cartesianMove(nextXMove, nextYMove)
	}
}

// DrawSandTunnelCurve draw a diamond around a center, with a bigger size than normal
func DrawSandTunnelCurve(img *image.RGBA64, points []Point, c color.RGBA64, sandCoef int) {
	for index, point := range points {
		if index+1 == len(points) {
			AddToSandTunnelLine(img, point, points[0], c, sandCoef)
		} else {
			AddToSandTunnelLine(img, point, points[index+1], c, sandCoef)
		}
	}
}

// Tunnel draws a turning shape
func Tunnel() {
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
	generatorPoints := tools.NewNumberGenerator(seed, 12, 32)
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

		color := tunnelCreateColor(i)

		DrawSandTunnelCurve(img, shape, color, param1)

		if updateImage != nil {
			updateImage(img)
		}
	}
}
