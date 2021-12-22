package model

import (
	"fmt"
	"flag"
)

type Arguments struct {
	Path string
	Deep int
}

func GetArguments() *Arguments {
	flag.Usage = func() {
		fmt.Printf("Usage: [options]\n Options:\n")
		flag.PrintDefaults()
	}

	path := flag.String("path", ".", "Path to list")	
	deep := flag.Int("deep", 0, "Depth to find")

	flag.Parse()

	return &Arguments{*path, *deep}
}
