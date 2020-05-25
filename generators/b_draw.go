package generators

import (
	"image"
	"image/color"
	"math"
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

func ipart(x float64) float64 {
	return math.Floor(x)
}

func round(x float64) float64 {
	return ipart(x + .5)
}

func fpart(x float64) float64 {
	return x - ipart(x)
}

func rfpart(x float64) float64 {
	return 1 - fpart(x)
}

// AddToWuLine draws line with Xialang Wu algorithm
func AddToWuLine(img *image.RGBA64, x1, y1, x2, y2 float64, w int, col color.RGBA64) {
	dx := x2 - x1
	dy := y2 - y1
	ax := dx
	if ax < 0 {
		ax = -ax
	}
	ay := dy
	if ay < 0 {
		ay = -ay
	}

	var plot func(int, int, float64)

	if ax < ay {
		x1, y1 = y1, x1
		x2, y2 = y2, x2
		dx, dy = dy, dx
		plot = func(x, y int, c float64) {
			newC := img.RGBA64At(y, x)

			if isLightBackground {
				if col.R <= newC.R {
					newC.R -= uint16(float64(col.R) * c)
				}

				if col.G <= newC.G {
					newC.G -= uint16(float64(col.G) * c)
				}

				if col.B <= newC.B {
					newC.B -= uint16(float64(col.B) * c)
				}
			} else {
				if maxValue-col.R >= newC.R {
					newC.R += uint16(float64(col.R) * c)
				}

				if maxValue-col.G >= newC.G {
					newC.G += uint16(float64(col.G) * c)
				}

				if maxValue-col.B >= newC.B {
					newC.B += uint16(float64(col.B) * c)
				}
			}

			img.SetRGBA64(y, x, newC)
		}
	} else {
		plot = func(x, y int, c float64) {
			newC := img.RGBA64At(x, y)

			if isLightBackground {
				if col.R <= newC.R {
					newC.R -= uint16(float64(col.R) * c)
				}

				if col.G <= newC.G {
					newC.G -= uint16(float64(col.G) * c)
				}

				if col.B <= newC.B {
					newC.B -= uint16(float64(col.B) * c)
				}
			} else {
				if maxValue-col.R >= newC.R {
					newC.R += uint16(float64(col.R) * c)
				}

				if maxValue-col.G >= newC.G {
					newC.G += uint16(float64(col.G) * c)
				}

				if maxValue-col.B >= newC.B {
					newC.B += uint16(float64(col.B) * c)
				}
			}

			img.SetRGBA64(x, y, newC)
		}
	}
	if x2 < x1 {
		x1, x2 = x2, x1
		y1, y2 = y2, y1
	}
	gradient := dy / dx

	xend := round(x1)
	yend := y1 + gradient*(xend-x1)
	xgap := rfpart(x1 + .5)
	xpxl1 := int(xend)
	ypxl1 := int(ipart(yend))
	plot(xpxl1, ypxl1, rfpart(yend)*xgap)
	plot(xpxl1, ypxl1+1, fpart(yend)*xgap)
	intery := yend + gradient

	xend = round(x2)
	yend = y2 + gradient*(xend-x2)
	xgap = fpart(x2 + 0.5)
	xpxl2 := int(xend)
	ypxl2 := int(ipart(yend))
	plot(xpxl2, ypxl2, rfpart(yend)*xgap)
	plot(xpxl2, ypxl2+1, fpart(yend)*xgap)

	for x := xpxl1 + 1; x <= xpxl2-1; x++ {
		plot(x, int(ipart(intery)), rfpart(intery))
		plot(x, int(ipart(intery))+1, fpart(intery))
		intery = intery + gradient
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

		pt := PolarPoint{
			angle:  p0.polar.angle,
			radius: p0.polar.radius + float64(sandCoef),
		}

		pr := pt.toCartesian(p0.center)

		//px := math.Cos(pt.angle)*(pt.radius) + float64(p0.center.x)
		//py := math.Sin(pt.angle)*(pt.radius) + float64(p0.center.y)

		AddToLine(img, p0.cartesian.x, p0.cartesian.y, pr.x, pr.y, c)

		//AddToWuLine(img, float64(p0.cartesian.x), float64(p0.cartesian.y), px, py, 1.0, c)

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
