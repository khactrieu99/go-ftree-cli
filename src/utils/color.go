package utils

import (
	"github.com/fatih/color"
)

func PrintWithColor(filename string, isDir bool) {
	if isDirs {
		color.Yellow(filename)
		return
	}
	color.White(filename)
}
