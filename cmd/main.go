package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mxgnus-de/discordbot-cli/internal/advancedbot"
	"github.com/mxgnus-de/discordbot-cli/internal/basebot"
	"github.com/mxgnus-de/discordbot-cli/internal/utils"
)

func main() {
	if len(os.Args) < 2 {
		ShowHelp()
		os.Exit(1)
	}

	utils.CheckIfNodeIsInstalled()
	utils.CheckIfNPMIsInstalled()

	basebotCmd := flag.NewFlagSet("basebot", flag.ExitOnError)
	typescriptBasebotPtr := basebotCmd.Bool("typescript", false, "Create a new bot with typescript")
	basebotCmd.Parse(os.Args[2:])
	advancedbotCmd := flag.NewFlagSet("advancedbot", flag.ExitOnError)
	typescriptAdvancedbotPtr := advancedbotCmd.Bool("typescript", false, "Create a new bot with typescript")
	advancedbotCmd.Parse(os.Args[2:])

	switch os.Args[1] {
	case "basebot":
		basebot.Create(basebotCmd, typescriptBasebotPtr)
	case "advancedbot":
		advancedbot.Create(advancedbotCmd, typescriptAdvancedbotPtr)
	default:
		ShowHelp()
		os.Exit(1)
	}
}

func ShowHelp() {
	fmt.Println("Usage: <command>")
	fmt.Println("Available commands:")
	fmt.Println("\tbasebot")
	fmt.Println("\tadvancedbot")
}
