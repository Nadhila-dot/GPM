package commands

import (
	nadhi "nadhi.dev/binaries/gpm/helpers"
	"nadhi.dev/binaries/gpm/toml"
)

func Setsource(newurl string) string {
	nadhi.Logo()
	// Removed cuz we are using the toml now ;)
	//configPath := "config.json"
	toml.EditConfigValueInToml("source", newurl)
	nadhi.Success("New package source set to " + newurl)

	return "true"

}


