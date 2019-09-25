package generators

import "image/color"

// Colors is a list of basic colors
var Colors = map[string]color.RGBA64{
	"empty": color.RGBA64{0x0000, 0x0000, 0x0000, 0x0000},
	"white": color.RGBA64{0xFFFF, 0xFFFF, 0xFFFF, 0xFFFF},
	"black": color.RGBA64{0x0000, 0x0000, 0x0000, 0xFFFF},
}

// ColorsPacks are group of colors
var ColorsPacks = map[string][]color.RGBA64{
	"trisummer": []color.RGBA64{
		color.RGBA64{0xF400, 0x4300, 0x3600, 0xFFFF},
		color.RGBA64{0xFF00, 0xC100, 0x0700, 0xFFFF},
		color.RGBA64{0xEF00, 0x6C00, 0x0000, 0xFFFF},
	},
	"green": []color.RGBA64{
		color.RGBA64{0x4C00, 0xAF00, 0x5000, 0xFFFF},
		color.RGBA64{0x3300, 0x6900, 0x1E00, 0xFFFF},
		color.RGBA64{0x9C00, 0xCC00, 0x6500, 0xFFFF},
		color.RGBA64{0x0000, 0x8900, 0x7B00, 0xFFFF},
	},
	"rainbow": []color.RGBA64{
		color.RGBA64{0xF400, 0x4300, 0x3600, 0xFFFF},
		color.RGBA64{0xFF00, 0x9800, 0x0000, 0xFFFF},
		color.RGBA64{0xFF00, 0xEB00, 0x3B00, 0xFFFF},
		color.RGBA64{0x4C00, 0xAF00, 0x5000, 0xFFFF},
		color.RGBA64{0x2100, 0x9600, 0xF300, 0xFFFF},
		color.RGBA64{0x3F00, 0x5100, 0xB500, 0xFFFF},
		color.RGBA64{0x9C00, 0x2700, 0xB000, 0xFFFF},
	},
	"beach": []color.RGBA64{
		color.RGBA64{0x2100, 0x9600, 0xF300, 0xFFFF},
		color.RGBA64{0xFF00, 0xC100, 0x0700, 0xFFFF},
	},
	"synthwave": []color.RGBA64{
		color.RGBA64{0x9200, 0x0000, 0x7500, 0xFFFF},
		color.RGBA64{0xFF00, 0x6C00, 0x1100, 0xFFFF},
		color.RGBA64{0x0D00, 0x0200, 0x2100, 0xFFFF},
	},
	"grey": []color.RGBA64{
		Colors["white"],
		Colors["black"],
	},
}

var emptyColor = Colors["empty"]
