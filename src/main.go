package main

import (
	"ftree/model"
)

func main() {
	args := model.GetArguments()
	model.GetFileTree(args.Path, args.Deep)
}
