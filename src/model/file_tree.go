package model

import (
	"fmt"
	"strings"
	"io/ioutil"
	"ftree/utils"
	"ftree/config"
)

func printSpaceChar(deep int) {
	for i:=1; i<=deep*config.SPACE_UNIT; i++ {
		fmt.Print(config.SPACE_CHAR)
	}
	fmt.Printf("%v ", config.FIRST_LINE_CHAR)
}

func loop(curPath string, curDeep int, maxDeep int) {
	if curDeep > maxDeep {
		return
	}

	files, err := ioutil.ReadDir(curPath)
	if err != nil {
		utils.HandleError(err)
	}
	
	for _, f := range files {

		// do nothing if file is hiding
		if strings.HasPrefix(f.Name(), ".") {
			continue
		}

		// print prefix space char
		printSpaceChar(curDeep)
	
		// print file name with color based on dir or normal file
		utils.PrintWithColor(f.Name(), f.IsDir())
		
		// loop child folder
		if f.IsDir() {
			childDirPath := curPath + "/" + f.Name()
			loop(childDirPath, curDeep+1, maxDeep)
		}
	}
}

func GetFileTree(path string, maxDeep int) {
	loop(path, 0, maxDeep)
}
