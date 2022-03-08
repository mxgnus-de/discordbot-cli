package advancedbot

import (
	"flag"
	"fmt"
	"os"

	"github.com/mxgnus-de/discordbot-cli/internal/utils"
	"gopkg.in/src-d/go-git.v4"
)

const tsBotGitURL = "https://github.com/mxgnus-de/discordjs-typescript-template.git"
const jsBotURL = "https://uploads.mxgnus.de/api/haste/695927be1f19267/raw"
const jsBotExampleCmdURL = "https://uploads.mxgnus.de/api/haste/8853787971819f2/raw"
const jsBotExampleEventURL = "https://uploads.mxgnus.de/api/haste/51369e16f6e6b3a/raw"

func Create(advancedbotcmd *flag.FlagSet, typescript *bool) {
	if *typescript {
		createWithTypescript()
	} else {
		createWithJavascript()
	}
}

func createWithTypescript() {
	utils.CheckIfGitIsInstalled()

	fmt.Println("Cloning typescript template...")
	_, err := git.PlainClone("./", false, &git.CloneOptions{
		URL: tsBotGitURL,
	})

	if err != nil {
		fmt.Println("Hmmm, something went wrong (GIT_CLONE_TYPESCRIPT_TEMPLATE)")
		os.Exit(1)
	}

	fmt.Println("Cloning typescript template done")
	fmt.Println("Installing dependencies...")
	utils.ExecCommand(
		"NPM_INSTALL_TYPESCRIPT_TEMPLATE",
		"npm", "install",
	)
	fmt.Println("Installing dependencies done")
	fmt.Println("------------------------------------------------------")
	fmt.Println("Setup done. Have fun with your bot!")
	fmt.Println("------------------------------------------------------")
}

func createWithJavascript() {
	fmt.Println("Creating a new bot...")
	fmt.Println("Setup npm...")
	utils.ExecCommand("NPM_INIT", "npm", "init", "-y")
	fmt.Println("Npm setup done.")

	fmt.Println("Install dependencie discord.js...")
	utils.ExecCommand("NPM_INSTALL_DISCORDJS", "npm", "install", "discord.js")
	fmt.Println("Dependencie discord.js installed.")

	fmt.Println("Get discord.js advancedbot...")
	content := utils.GetDataFromURL("JAVASCRIPT_GET_advancedbot", jsBotURL)
	fmt.Println("Get discord.js advancedbot done.")

	fmt.Println("Creating index.js...")
	utils.CreateFile("JAVASCRIPT_CREATE_INDEX", "index.js", content)
	fmt.Println("index.js created.")

	fmt.Println("Create commands folder...")
	utils.CreateFolder("JAVASCRIPT_CREATE_COMMANDS_FOLDER", "commands")
	fmt.Println("commands folder created.")

	fmt.Println("Get example command...")
	content = utils.GetDataFromURL("JAVASCRIPT_GET_EXAMPLE_COMMAND", jsBotExampleCmdURL)
	fmt.Println("Get example command done.")

	fmt.Println("Create example command...")
	utils.CreateFile("JAVASCRIPT_CREATE_EXAMPLE_COMMAND", "commands/example.js", content)
	fmt.Println("example command created.")

	fmt.Println("Create events folder...")
	utils.CreateFolder("JAVASCRIPT_CREATE_EVENTS_FOLDER", "events")
	fmt.Println("events folder created.")

	fmt.Println("Get example event...")
	content = utils.GetDataFromURL("JAVASCRIPT_GET_EXAMPLE_EVENT", jsBotExampleEventURL)
	fmt.Println("Get example event done.")

	fmt.Println("Create example event...")
	utils.CreateFile("JAVASCRIPT_CREATE_EXAMPLE_EVENT", "events/example.js", content)
	fmt.Println("example event created.")

	fmt.Println("------------------------------------------------------")
	fmt.Println("Setup done. Have fun with your bot!")
	fmt.Println("------------------------------------------------------")
}
