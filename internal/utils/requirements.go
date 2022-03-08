package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func CheckIfNodeIsInstalled() {
	cmd := exec.Command("node", "--version")
	_, err := cmd.Output()

	if err != nil {
		fmt.Println("Node is not installed. Please install node first. (https://nodejs.org/en/download/)")
		os.Exit(1)
	}
}

func CheckIfNPMIsInstalled() {
	cmd := exec.Command("npm", "--version")
	_, err := cmd.Output()

	if err != nil {
		fmt.Println("NPM is not installed. Please install npm first. (https://www.npmjs.com/get-npm)")
		os.Exit(1)
	}
}

func CheckIfGitIsInstalled() {
	cmd := exec.Command("git", "--version")
	_, err := cmd.Output()

	if err != nil {
		fmt.Println("Git is not installed. Please install git first. (https://git-scm.com/downloads)")
		os.Exit(1)
	}
}
