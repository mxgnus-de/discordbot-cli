package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mxgnus-de/discordbot-cli/internal/basebot"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: basebot <command>")
		os.Exit(1)
	}

	checkIfNodeIsInstalled()
	checkIfNPMIsInstalled()

	basebotCmd := flag.NewFlagSet("basebot", flag.ExitOnError)
	typescriptPtr := basebotCmd.Bool("typescript", false, "Create a new bot with typescript")
	basebotCmd.Parse(os.Args[2:])

	switch os.Args[1] {
	case "basebot":
		basebot.Create(basebotCmd, typescriptPtr)
	default:
		fmt.Println("Usage: <command>")
		fmt.Println("Available commands:")
		fmt.Println("\tbasebot")
		os.Exit(1)
	}
}

func checkIfNodeIsInstalled() {
	if _, err := os.Stat("node"); os.IsNotExist(err) {
		fmt.Println("Node is not installed. Please install node first. (https://nodejs.org/en/download/)")
		os.Exit(1)
	}
}

func checkIfNPMIsInstalled() {
	if _, err := os.Stat("npm"); os.IsNotExist(err) {
		fmt.Println("NPM is not installed. Please install npm first. (https://www.npmjs.com/get-npm)")
		os.Exit(1)
	}
}
