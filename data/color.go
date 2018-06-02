package data

type Color struct {
	Color     int
	Name      string
	ShortName string
	Debug     string
}

func (c Color) String() string {
	if !GetDebug() {
		return c.ShortName
	} else {
		return c.Debug
	}
}
