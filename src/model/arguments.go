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
	path := args.Path
	file, err := os.Stat(path)
	if os.IsNotExist(err) {
		return errors.New("argsErr: Path does not exists")
	}
	if !file.Mode().IsDir() {
		return errors.New("argsErr: Destination must be a folder")
	}

	deep := args.Deep
	if deep<0 {
		return errors.New("argsErr: Depth must not be negative")
	}

	return nil
}

func GetArguments() *Arguments {
	flag.Usage = func() {
		fmt.Printf("Usage: [options]\n Options:\n")
		flag.PrintDefaults()
	}

	path := flag.String("path", ".", "Path to list")	
	deep := flag.Int("deep", 0, "Depth to find")

	flag.Parse()

	args := &Arguments{*path, *deep}
	err := validate(args)
	if err != nil {
		utils.HandleError(err)
	}

	return args
}
