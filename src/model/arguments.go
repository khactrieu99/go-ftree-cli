package model

import (
	"fmt"
	"flag"
	"os"
	"errors"
	"ftree/utils"
)

type Arguments struct {
	Path string
	Deep int
}

func validate(args *Arguments) error {

	// validate path
	path := args.Path
	file, err := os.Stat(path)

	if os.IsNotExist(err) {
		return errors.New("argsErr: Path does not exists")
	}

	if !file.Mode().IsDir() {
		return errors.New("argsErr: Destination must be a folder")
	}

	// validate deep
	deep := args.Deep
	if deep<0 {
		return errors.New("argsErr: Depth must not be negative")
	}

	return nil
}

func GetArguments() *Arguments {

	// use for flag -h (help)
	flag.Usage = func() {
		fmt.Printf("Usage: [options]\n Options:\n")
		flag.PrintDefaults()
	}

	// define flag
	path := flag.String("path", ".", "Path to list")	
	deep := flag.Int("deep", 0, "Depth to find")

	// get input
	flag.Parse()

	// validate provided arguments
	args := &Arguments{*path, *deep}
	err := validate(args)
	if err != nil {
		utils.HandleError(err)
	}

	return args
}
