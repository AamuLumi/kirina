package generators

import (
	"image"
	"image/color"
)

var maxValue = uint16(0xFFFF)
var minValue = uint16(0x0000)

// Line draws line by Bresenham's algorithm.
func Line(img *image.RGBA64, x0, y0, x1, y1 int, c color.RGBA64) {
	// implemented straight from WP pseudocode
	dx := x1 - x0
	if dx < 0 {
		dx = -dx
	}
	dy := y1 - y0
	if dy < 0 {
		dy = -dy
	}
	var sx, sy int
	if x0 < x1 {
		sx = 1
	} else {
		sx = -1
	}
	if y0 < y1 {
		sy = 1
	} else {
		sy = -1
	}
	err := dx - dy

	for {
		img.Set(x0, y0, c)
		if x0 == x1 && y0 == y1 {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}

// AddToLine draws line by Bresenham's algorithm.
func AddToLine(img *image.RGBA64, x0, y0, x1, y1 int, c color.RGBA64) {
	// implemented straight from WP pseudocode
	dx := x1 - x0
	if dx < 0 {
		dx = -dx
	}
	dy := y1 - y0
	if dy < 0 {
		dy = -dy
	}
	var sx, sy int
	if x0 < x1 {
		sx = 1
	} else {
		sx = -1
	}
	if y0 < y1 {
		sy = 1
	} else {
		sy = -1
	}
	err := dx - dy

	for {
		newC := img.RGBA64At(x0, y0)

		if isLightBackground {
			if c.R <= newC.R {
				newC.R -= c.R
			}

			if c.G <= newC.G {
				newC.G -= c.G
			}

			if c.B <= newC.B {
				newC.B -= c.B
			}
		} else {
			if maxValue-c.R >= newC.R {
				newC.R += c.R
			}

			if maxValue-c.G >= newC.G {
				newC.G += c.G
			}

			if maxValue-c.B >= newC.B {
				newC.B += c.B
			}
		}

		img.SetRGBA64(x0, y0, newC)

		if x0 == x1 && y0 == y1 {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}

// AddToDegressiveLine draws line by Bresenham's algorithm.
func AddToDegressiveLine(img *image.RGBA64, x0, y0, x1, y1 int, c color.RGBA64) {
	totalDistance := distance(x0, y0, x1, y1)

	// implemented straight from WP pseudocode
	dx := x1 - x0
	if dx < 0 {
		dx = -dx
	}
	dy := y1 - y0
	if dy < 0 {
		dy = -dy
	}
	var sx, sy int
	if x0 < x1 {
		sx = 1
	} else {
		sx = -1
	}
	if y0 < y1 {
		sy = 1
	} else {
		sy = -1
	}
	err := dx - dy

	for {
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}

		newC := img.RGBA64At(x0, y0)
		currentDistance := distance(x0, y0, x1, y1)

		rValue := uint16(float64(c.R) * (totalDistance - currentDistance) / totalDistance)
		gValue := uint16(float64(c.G) * (totalDistance - currentDistance) / totalDistance)
		bValue := uint16(float64(c.B) * (totalDistance - currentDistance) / totalDistance)

		if isLightBackground {
			if rValue <= newC.R {
				newC.R -= rValue
			}

			if gValue <= newC.G {
				newC.G -= gValue
			}

			if bValue <= newC.B {
				newC.B -= bValue
			}
		} else {
			if maxValue-rValue >= newC.R {
				newC.R += rValue
			}

			if maxValue-gValue >= newC.G {
				newC.G += gValue
			}

			if maxValue-bValue >= newC.B {
				newC.B += bValue
			}
		}

		img.SetRGBA64(x0, y0, newC)

		if x0 == x1 && y0 == y1 {
			break
		}
	}
}

/**
// AddToSandLine draws line by Bresenham's algorithm.
func AddToSandLine(img *image.RGBA64, p0, p1 AngularPoint, c color.RGBA64, sandCoef int) {
	sandX := int(math.Cos(p0.angle) * (p0.currentDistance + float64(sandCoef)))
	sandY := int(math.Sin(p0.angle) * (p0.currentDistance + float64(sandCoef)))

	// implemented straight from WP pseudocode
	dx := p1.x - p0.x
	if dx < 0 {
		dx = -dx
	}

	dy := p1.y - p0.y
	if dy < 0 {
		dy = -dy
	}

	var sx, sy int
	if p0.x < p1.x {
		sx = 1
	} else {
		sx = -1
	}

	if p0.y < p1.y {
		sy = 1
	} else {
		sy = -1
	}

	err := dx - dy

	for {
		AddToDegressiveLine(img, p0.x, p0.y, sandX+p0.x, sandY+p0.y, c)

		if p0.x == p1.x && p0.y == p1.y {
			break
		}

		e2 := 2 * err

		if e2 > -dy {
			err -= dy
			p0.x += sx
		}

		if e2 < dx {
			err += dx
			p0.y += sy
		}
	}
}
*/

// AddToSandLine draws line by Bresenham's algorithm.
func AddToSandLine(img *image.RGBA64, p0, p1 Point, c color.RGBA64, sandCoef int) {
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

		AddToDegressiveLine(img, p0.cartesian.x, p0.cartesian.y, pr.x, pr.y, c)

		if p0.cartesian.x == p1.cartesian.x && p0.cartesian.y == p1.cartesian.y {
			break
		}
	}
}

// DiamondBigger draw a diamond around a center, with a bigger size than normal
func DiamondBigger(img *image.RGBA64, cx, cy, size int, c color.RGBA64) {
	bounds := img.Bounds()

	for y := cy - size; y < cy+size; y++ {
		for x := cx - size; x < cx+size; x++ {
			if x >= 0 && x <= bounds.Max.X && y >= 0 && y <= bounds.Max.Y && abs(cy-y)+abs(cx-x) <= size+1 {
				img.Set(x, y, c)
			}
		}
	}
}

// DrawSandCurve draw a diamond around a center, with a bigger size than normal
func DrawSandCurve(img *image.RGBA64, points []Point, c color.RGBA64, sandCoef int) {
	for index, point := range points {
		if index+1 == len(points) {
			AddToSandLine(img, point, points[0], c, sandCoef)
		} else {
			AddToSandLine(img, point, points[index+1], c, sandCoef)
		}
	}
}

// DrawPoint just draws a point of a color
func DrawPoint(img *image.RGBA64, cx, cy int, c color.RGBA64) {
	img.Set(cx, cy, c)
}

// Round do a round
func Round(img *image.RGBA64, x0, y0, x1, y1 int, c color.RGBA64) {

}

// ImageOpacity change opacity of drawn pixels
func ImageOpacity(img *image.RGBA64, c color.RGBA64) {
	bounds := img.Bounds()

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			currentColor := img.RGBA64At(x, y)

			currentColor.R = uint16((currentColor.R + c.R))
			currentColor.G = uint16((currentColor.G + c.G))
			currentColor.B = uint16((currentColor.B + c.B))

			img.SetRGBA64(x, y, currentColor)
		}
	}
}
