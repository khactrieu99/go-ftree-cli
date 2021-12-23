package main

import (
	"fmt"
	"ftree/model"
)

func main() {
	args := model.GetArguments()
	model.GetFileTree(args.Path, args.Deep)
}
