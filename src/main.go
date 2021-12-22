package main

import (
	"fmt"
	"ftree/model"
)

func main() {
	args := model.GetArguments()
	fmt.Println(args.Path)
}
