package main

import (
	"fmt"
	"os"

	nadhi "nadhi.dev/binaries/gpm/helpers"
)

func main() {
	if len(os.Args) < 2 {
        nadhi.Logo()
		fmt.Println("Usage: gpm <command> [parameters] \nDo gpm help for more info.")
		return
	}

	nadhi.CreateConfigFile("config.json")

	handleCommand(os.Args[1:])
}
