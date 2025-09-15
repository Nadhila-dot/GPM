package main

import (
	"fmt"
	"os"

	nadhi "nadhi.dev/binaries/gpm/helpers"
	"nadhi.dev/binaries/gpm/toml"
)

func main() {

    if !nadhi.HasGoMod() {
        nadhi.Hint("No go.mod file found. To install packages, please run 'go mod init' first.")
		
	}

    // This is important
    toml.ParseConfigFromToml()

	if len(os.Args) < 2 {
        nadhi.Logo()
		fmt.Println("Usage: gpm <command> [parameters] \nDo gpm help for more info.")
		return
	}

    // Deprecated and moved to the toml file
	//nadhi.CreateConfigFile("config.json")

	handleCommand(os.Args[1:])
}
