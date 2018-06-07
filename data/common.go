package data

var (
	Debug = false
)

const (
	White  = "#FFFFFF"
	Yellow = "#FFD500"
	Red    = "#C41E3A"
	Orange = "#FF5800"
	Blue   = "#0051BA"
	Green  = "#009E60"
	Gray   = "#383838"
)

func ToggleDebug() {
	SetDebug(!GetDebug())
}

func GetDebug() bool {
	return Debug
}

func SetDebug(b bool) {
	Debug = b
}

func GetColorsOLL(color string) map[string]Color {
	c := make(map[string]Color, len(Colors))
	for k, v := range Colors {
		c[k] = v
		if k != color {
			val := c[k]
			val.HexColor = Gray
			c[k] = val
		}
	}

	return c
}

// Default colors
var (
	Colors = map[string]Color{
		"white": Color{
			Color:     0,
			Name:      "white",
			ShortName: "W",
			HexColor:  White,
		},
		"orange": Color{
			Color:     1,
			Name:      "orange",
			ShortName: "O",
			HexColor:  Orange,
		},
		"green": Color{
			Color:     2,
			Name:      "green",
			ShortName: "G",
			HexColor:  Green,
		},
		"red": Color{
			Color:     3,
			Name:      "red",
			ShortName: "R",
			HexColor:  Red,
		},
		"blue": Color{
			Color:     4,
			Name:      "blue",
			ShortName: "B",
			HexColor:  Blue,
		},
		"yellow": Color{
			Color:     5,
			Name:      "yellow",
			ShortName: "Y",
			HexColor:  Yellow,
		},
	}
)
