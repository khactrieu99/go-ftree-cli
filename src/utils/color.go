package utils

import (
	"github.com/fatih/color"
)

func PrintWithColor(filename string, isDir bool) {
	if isDir {
		color.Cyan(filename)
		return
	}
	color.Green(filename)
}
