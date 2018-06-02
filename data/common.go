package data

var (
	Debug = false
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

// Default colors
var (
	Colors = map[string]Color{
		"white":  Color{0, "white", "W", ""},
		"orange": Color{1, "orange", "O", ""},
		"green":  Color{2, "green", "G", ""},
		"red":    Color{3, "red", "R", ""},
		"blue":   Color{4, "blue", "B", ""},
		"yellow": Color{5, "yellow", "Y", ""},
	}
)
