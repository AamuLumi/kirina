package generators

import (
	"image"
	"image/color"

	"kirina/tools"
)

var baseColor color.RGBA64
var seed int
var img *image.RGBA64
var updateImage func(*image.RGBA64)
var param1 = -1
var param2 = -1
var param3 = -1
var cycles = -1
var colors []color.RGBA64
var isLightBackground = false

// SetBaseColor sets the base color of image
func SetBaseColor(color color.RGBA64) {
	baseColor = color

	if color.R > 0x8000 && color.G > 0x8000 && color.B > 0x8000 {
		isLightBackground = true

		if len(colors) > 0 {
			colors = tools.ColorsInversion(colors)
		}
	}
}

// SetNumberSeed sets the seed for number generator
func SetNumberSeed(s int) {
	seed = s
}

// SetImage sets the target image
func SetImage(i *image.RGBA64) {
	img = i
}

// SetDrawer sets drawing function
func SetDrawer(f func(*image.RGBA64)) {
	updateImage = f
}

// SetParam1 sets the first parameter for generator
func SetParam1(p int) {
	param1 = p
}

// SetParam2 sets the second parameter for generator
func SetParam2(p int) {
	param2 = p
}

// SetParam3 sets the third parameter for generator
func SetParam3(p int) {
	param3 = p
}

// SetParamColors sets the colors parameter for generator
func SetParamColors(c []color.RGBA64) {
	colors = c

	if isLightBackground {
		colors = tools.ColorsInversion(c)
	}
}

// SetCycles sets number of cycles of the generator to do
func SetCycles(nb int) {
	cycles = nb
}
