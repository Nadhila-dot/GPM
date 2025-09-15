package commands

import (
	nadhi "nadhi.dev/binaries/gpm/helpers"
	"nadhi.dev/binaries/gpm/toml"
)

func Remove(packageName string) string {
	nadhi.Logo()
	if !nadhi.HasGoMod() {
		nadhi.Error("No go.mod file found. Please run 'go mod init' first.")
		return "Error: no go.mod file found"
	}
	result := toml.CheckToml()
	if !result {
		nadhi.Error("gpm.toml file does not exist.. Did you install any package?")
		return "no toml"
	}

	nadhi.LoadingCheck("Getting package info for " + packageName + "..")
	path := toml.GetImportPath(packageName)

	if path == "" {
		nadhi.Error("Package not found: " + packageName)
		return "Package not found: " + packageName
	}

	nadhi.LoadingCheck("Starting uninstallation of " + packageName + "..")
	nadhi.Uninstall(path)

	nadhi.LoadingCheck("Removing " + packageName + " from gpm.toml file..")
	resultToml := toml.TomlRemove(packageName)
	if resultToml != "success" {
		nadhi.Error(resultToml)
		return "Error: " + resultToml
	}
	nadhi.Success("Successfully removed " + packageName)
	nadhi.VerifyPackagesAndRemoveOnError()

	// Keep this here, we want to show this after everything is done
	return "Removed package: " + packageName
}
