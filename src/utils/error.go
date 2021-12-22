package ftree

import (
	"fmt"
	"os"
)

func HandleError(err *Error) {
	fmt.Print("")
	os.Exit()
}
