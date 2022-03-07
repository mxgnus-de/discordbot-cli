package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

func CreateFile(module string, path string, content []byte) {
	err := ioutil.WriteFile(path, content, 0644)

	if err != nil {
		fmt.Println("Hmmm, something went wrong (" + module + ")")
		os.Exit(1)
	}
}
