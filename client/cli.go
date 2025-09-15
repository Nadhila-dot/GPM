package main

import (
	"fmt"

	"nadhi.dev/binaries/gpm/commands"
	nadhi "nadhi.dev/binaries/gpm/helpers"
)

// This server handles the CLI commands
// Takes command and calls function, sounds simple enough, maybe later we can add flags too

func handleCommand(args []string) {
	if len(args) < 1 {
		fmt.Println("No command provided.")
		return
	}

	command := args[0]

	switch command {

	case "install":
		if !hasEnoughArgs(args, 2, "install") {
			return
		}
		param := args[1]
		commands.Download(param)
    case "build":
        if !hasEnoughArgs(args, 2, "build") {
            return
        }
        appname := args[1]
        dir := "./"
        ignoreos := []string{}
        if len(args) > 2 {
            // Any additional args after appname are OSes to ignore
            ignoreos = args[2:]
        }
        nadhi.Build(appname, dir, ignoreos)
    case "tidy":
		if !hasEnoughArgs(args, 1, "install") {
			return
		}
		
		commands.Tidy()
	case "add":
		if !hasEnoughArgs(args, 2, "add") {
			return
		}
		param := args[1]
		commands.Download(param)
	
	case "remove":
		if !hasEnoughArgs(args, 2, "remove") {
			return
		}
		param := args[1]
		commands.Remove(param)
	case "refersh":
		if len(args) > 1 {
			commands.Refresh(args[1])
		} else {
			commands.Refresh("")
		}

	case "help":
		if !hasEnoughArgs(args, 1, "help") {
			return
		}
		commands.Help()
	case "list":
		if !hasEnoughArgs(args, 1, "list") {
			return
		}

		commands.List()
	case "setsource":
		if !hasEnoughArgs(args, 2, "setsource") {
			return
		}

		commands.Setsource(args[1])
	default:

		nadhi.Error("Unknown command: " + command + "\nDo gpm help to see available commands.")
	}
}

func hasEnoughArgs(args []string, needed int, command string) bool {
	if len(args) < needed {
		fmt.Printf("Usage: gpm %s <parameter>\n", command)
		return false
	}
	return true
}
