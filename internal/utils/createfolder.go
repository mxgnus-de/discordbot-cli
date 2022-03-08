package utils

import (
	"fmt"
	"os"
)

func CreateFolder(module string, folder string) {
	err := os.Mkdir(folder, os.ModePerm)
	if err != nil {
		fmt.Println("Hmmm, something went wrong (" + module + ")")
		os.Exit(1)
	}
}
