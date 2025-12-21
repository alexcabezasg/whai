package ui

import (
	"strings"

	"github.com/pterm/pterm"
)

type DefaultUI struct{}

var colorMap = map[Color]pterm.Color{
	White:  pterm.FgWhite,
	Yellow: pterm.FgYellow,
	Green:  pterm.FgGreen,
	Cyan:   pterm.FgCyan,
}

func (ui DefaultUI) Print(str string, format Format) {
	getStyle(format).Print(strings.Repeat(" ", format.Indentation) + str)
}

func (ui DefaultUI) Println(str string, format Format) {
	getStyle(format).Println(strings.Repeat(" ", format.Indentation) + str)
}

func (ui DefaultUI) EmptyLine() {
	pterm.Println()
}

func getStyle(format Format) *pterm.Style {
	color := getColor(format.Color)
	if format.Bold {
		return pterm.NewStyle(color, pterm.Bold)
	}
	return pterm.NewStyle(color)
}

func getColor(color Color) pterm.Color {
	return colorMap[color]
}
