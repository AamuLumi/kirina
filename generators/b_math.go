package generators

import (
	"math"
)

func abs(x int) int {
	if x < 0 {
		return x * -1
	}

	return x
}

func distance(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)))
}

// Point is a mathematical point
type Point struct {
	polar     PolarPoint
	cartesian CartesianPoint
	center    CartesianPoint
}

func (p *Point) cartesianMove(dx, dy int) {
	p.cartesian.x += dx
	p.cartesian.y += dy

	p.polar = p.cartesian.toPolar(p.center)
}

func (p *Point) polarMove(da, dr float64) {
	p.polar.radius += dr
	p.polar.angle += da

	p.cartesian = p.polar.toCartesian(p.center)
}

// PolarPoint is an polar representation of point
type PolarPoint struct {
	angle  float64
	radius float64
}

func (p *PolarPoint) toCartesian(center CartesianPoint) CartesianPoint {
	return CartesianPoint{
		x: int(math.Cos(p.angle)*(p.radius)) + center.x,
		y: int(math.Sin(p.angle)*(p.radius)) + center.y,
	}
}

// CartesianPoint is a cartesian representation of point
type CartesianPoint struct {
	x int
	y int
}

func (p *CartesianPoint) toPolar(center CartesianPoint) PolarPoint {
	opposite := float64(p.y - center.y)
	adjacent := float64(p.x - center.x)
	dist := distance(p.x, p.y, center.x, center.y)

	s := math.Asin(opposite / dist)
	c := math.Acos(adjacent / dist)

	if s < 0 {
		c = c * -1
	}

	return PolarPoint{
		angle:  c,
		radius: distance(p.x, p.y, center.x, center.y),
	}
}
