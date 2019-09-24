package tools

import (
	"image/color"
)

// ColorGenerator is a color generator
type ColorGenerator struct {
	currentValue    int
	maxColors       []color.RGBA64
	seed            int
	increasingValue bool
	nbColors        int
}

// NewColorGenerator create a color generator
func NewColorGenerator(seed int, maxColors []color.RGBA64) ColorGenerator {
	return ColorGenerator{
		currentValue:    0,
		maxColors:       maxColors,
		seed:            seed,
		increasingValue: true,
		nbColors:        len(maxColors) * 256,
	}
}

func mean(percentage int, a, b uint16) uint16 {
	return uint16(float64(percentage)/100*float64(b) + float64(100-percentage)/100*float64(a))
}

func computeMeanOf(percentage int, min, max *color.RGBA64) color.RGBA64 {
	r := mean(percentage, min.R, max.R)
	g := mean(percentage, min.G, max.G)
	b := mean(percentage, min.B, max.B)
	a := mean(percentage, min.A, max.A)

	return color.RGBA64{r, g, b, a}
}

// ColorMean do the mean between two colors for a given percentage
func ColorMean(percentage int, min, max *color.RGBA64) color.RGBA64 {
	return computeMeanOf(percentage, min, max)
}

// NextColor returns the next color of the generator
func (g *ColorGenerator) NextColor() color.RGBA64 {
	maxColorToGet := g.currentValue / 256

	var maxColorDown, maxColorUp color.RGBA64

	if maxColorToGet+1 == len(g.maxColors) {
		maxColorDown, maxColorUp = g.maxColors[maxColorToGet], g.maxColors[0]
	} else {
		maxColorDown, maxColorUp = g.maxColors[maxColorToGet], g.maxColors[maxColorToGet+1]
	}

	computedColor := computeMeanOf(g.currentValue%256*100/256, &maxColorDown, &maxColorUp)

	g.currentValue = (g.currentValue + 1) % g.nbColors

	return computedColor
}
