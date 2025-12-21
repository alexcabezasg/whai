package ui

type Format struct {
	Color       Color
	Bold        bool
	Indentation int
}

type Color int

const (
	White Color = iota
	Cyan
	Yellow
	Green
)
