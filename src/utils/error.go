package utils

import (
	"fmt"
	"os"
)

func HandleError(e error) {
	fmt.Fprintf(os.Stderr, "%v\n", e.Error())
	os.Exit(-1)
}
