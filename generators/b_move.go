package generators

import (
	"../tools"
)

func basic4MoveWithCoef(x, y *int, coef int, g *tools.NumberGenerator) {
	switch g.NextPositive() {
	case 3:
		*x += coef
		*y += coef
	case 2:
		*x += coef
		*y -= coef
	case 1:
		*x -= coef
		*y += coef
	default:
		*x -= coef
		*y -= coef
	}
}

func basic8MoveWithCoef(x, y *int, coef int, g *tools.NumberGenerator) {
	switch g.NextPositive() {
	case 7:
		*x += coef / 2
	case 6:
		*x -= coef / 2
	case 5:
		*y += coef / 2
	case 4:
		*y -= coef / 2
	case 3:
		*x += coef
		*y += coef
	case 2:
		*x += coef
		*y -= coef
	case 1:
		*x -= coef
		*y += coef
	default:
		*x -= coef
		*y -= coef
	}
}

func moveAwayFromCenterWithCoef(cx, cy int, point *Point, coef int, g100, g4 *tools.NumberGenerator) {
	//randomNumberX := g.NextPositive()
	//randomNumberY := g.NextPositive()

	point.polarMove(float64(g100.NextPositive()-50)/3000, float64(g4.NextPositive()))

	// if *x < cx {
	// 	if abs(*x-cx) > 200 && randomNumberX < 3 {
	// 		*x += coef * randomNumberX
	// 	} else {
	// 		*x -= coef * randomNumberX
	// 	}
	// } else if *x > cx {
	// 	if abs(*x-cx) > 200 && randomNumberX < 3 {
	// 		*x -= coef * randomNumberX
	// 	} else {
	// 		*x += coef * randomNumberX
	// 	}
	// }

	// if *y < cy {
	// 	if abs(*y-cy) > 200 && randomNumberY < 3 {
	// 		*y += coef * randomNumberY
	// 	} else {
	// 		*y -= coef * randomNumberY
	// 	}
	// } else if *y > cy {
	// 	if abs(*y-cy) > 200 && randomNumberY < 3 {
	// 		*y -= coef * randomNumberY
	// 	} else {
	// 		*y += coef * randomNumberY
	// 	}
	// }
}

func basicMoveLine(x, y, x2, y2 *int, coef int, g *tools.NumberGenerator) {
	switch g.NextPositive() {
	case 7:
		*x += coef
		*x2 += coef
	case 6:
		*x -= coef
		*x -= coef
	case 5:
		*y += coef
		*y += coef
	case 4:
		*y -= coef
		*y -= coef
	case 3:
		*x += coef
		*x2 += coef
		*y += coef
		*y2 += coef
	case 2:
		*x += coef
		*x2 += coef
		*y -= coef
		*y2 -= coef
	case 1:
		*x -= coef
		*x2 -= coef
		*y += coef
		*y2 += coef
	default:
		*x -= coef
		*x2 -= coef
		*y -= coef
		*y2 -= coef
	}
}
