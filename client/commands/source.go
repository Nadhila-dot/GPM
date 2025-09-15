package commands

import nadhi "nadhi.dev/binaries/gpm/helpers"

func Setsource(newurl string) string {
	nadhi.Logo()
	configPath := "config.json"
	nadhi.EditConfigValue(configPath, "source", newurl)
	nadhi.Success("New package source set to " + newurl)

	return "true"

}


