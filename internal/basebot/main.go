package basebot

import (
	"flag"
	"fmt"
	"os"

	"github.com/mxgnus-de/discordbot-cli/internal/utils"
)

var jsBaseBotURL = "https://uploads.mxgnus.de/api/haste/1517964bf73ef35/raw"
var tsBaseBotURL = "https://uploads.mxgnus.de/api/haste/246b6c97b8907d0/raw"

func Create(basebotCmd *flag.FlagSet, typescript *bool) {
	if *typescript {
		createWithTypescript()
	} else {
		createWithJavascript()
	}
}

func createWithJavascript() {

	fmt.Println("Creating a new bot...")
	fmt.Println("Setup npm...")
	utils.ExecCommand("NPM_INIT", "npm", "init", "-y")
	fmt.Println("Npm setup done.")

	fmt.Println("Install dependencie discord.js...")
	utils.ExecCommand("NPM_INSTALL_DISCORDJS", "npm", "install", "discord.js")
	fmt.Println("Dependencie discord.js installed.")

	fmt.Println("Get discord.js basebot...")
	content := utils.GetDataFromURL("JAVASCRIPT_GET_BASEBOT", jsBaseBotURL)
	fmt.Println("Get discord.js basebot done.")

	fmt.Println("Creating index.js...")
	utils.CreateFile("JAVASCRIPT_CREATE_INDEX", "index.js", content)
	fmt.Println("index.js created.")

	fmt.Println("------------------------------------------------------")
	fmt.Println("Setup done. Have fun with your bot!")
	fmt.Println("------------------------------------------------------")
}

func createWithTypescript() {
	fmt.Println("Creating a new bot...")
	fmt.Println("Setup npm...")
	utils.ExecCommand("NPM_INIT", "npm", "init", "-y")
	fmt.Println("Npm setup done.")

	fmt.Println("Install dependencie discord.js...")
	utils.ExecCommand("NPM_INSTALL_DISCORDJS", "npm", "install", "discord.js")
	fmt.Println("Dependencie discord.js installed.")

	fmt.Println("Install typescript...")
	utils.ExecCommand("NPM_INSTALL_TYPESCRIPT", "npm", "install", "typescript", "-D")
	fmt.Println("Typescript installed.")

	fmt.Println("Setup typescript...")
	utils.ExecCommand("TYPESCRIPT_SETUP", "npx", "tsc", "--init", "--outDir", "./dist", "--rootDir", "./src")
	fmt.Println("Setup typescript done.")

	fmt.Println("Creating src directory...")
	err := os.Mkdir("src", os.ModePerm)

	if err != nil {
		fmt.Println("Error creating src directory: " + err.Error())
		os.Exit(1)
	}

	fmt.Println("Get discord.js typescript basebot...")
	content := utils.GetDataFromURL("TYPESCRIPT_GET_BASEBOT", tsBaseBotURL)
	fmt.Println("Get discord.js typescript basebot done.")

	fmt.Println("Creating src/index.ts...")
	utils.CreateFile("TYPESCRIPT_CREATE_INDEX", "src/index.ts", content)
	fmt.Println("src/index.ts created.")

	fmt.Println("------------------------------------------------------")
	fmt.Println("Setup typescript done. Have fun with your bot!")
	fmt.Println("------------------------------------------------------")
}
