package commands

import (
    "fmt"
    "strings"

    "github.com/fatih/color"
    nadhi "nadhi.dev/binaries/gpm/helpers"
    "nadhi.dev/binaries/gpm/toml"
)

func Download(packageArg string) string {
	nadhi.Logo()
	if !nadhi.HasGoMod() {
		nadhi.Error("No go.mod file found. Please run 'go mod init' first.")
		return "Error: no go.mod file found"
	}
    result := toml.CreateGpmFile()
    if result != "success" {
        return "Error: " + result
    }

    nadhi.SuccessCheck("Found gpm.toml file..")

    source, err := nadhi.GetConfig("source")
    if err != nil {
        nadhi.Error(fmt.Sprintf("Error getting config: %v", err))
        return ""
    }
    nadhi.SuccessCheck(fmt.Sprintf("Using package source: %s", source.Source))

    pkgName := packageArg
    version := ""
    if strings.Contains(packageArg, "@") {
        parts := strings.SplitN(packageArg, "@", 2)
        pkgName = parts[0]
        version = parts[1]
    }

    pkgData, err := nadhi.SearchPackages(source.Source, pkgName)
    nadhi.SuccessCheck("Fetching packages from source..")
    if err != nil {
        if nadhi.CheckString([]string{"invalid character"}, err.Error()) {
            nadhi.Error(color.New(color.BgHiGreen).Sprint("Current Source is either unavailable or invalid"))
            return "Current Source is either unavailable or invalid"
        }
        nadhi.Error(fmt.Sprintf("Error fetching package info: %v", err))
        return ""
    }

    pkg, ok := pkgData.Packages[pkgName]
    if !ok {
        nadhi.Error("Package not found: " + pkgName)
        return "Error: package not found"
    }

    versions := pkg.(map[string]interface{})["versions"].(map[string]interface{})
    importPath := ""

    if version == "" || version == "latest" {
        importPath = versions["default"].(string)
    } else if major, ok := versions["major"].(map[string]interface{}); ok {
        verParts := strings.SplitN(version, ".", 2)
        majorKey := verParts[0]
        if v, ok := major[majorKey]; ok {
            importPath = v.(string)
            if len(verParts) > 1 {
                version = verParts[1]
            } else {
                version = ""
            }
        }
    }

    if importPath == "" {
        nadhi.Error("Version not found for package: " + pkgName)
        return "Error: version not found"
    }

    nadhi.Install(source.Source, packageArg)

    resultToml := toml.TomlAdd(pkgName, importPath)
    if resultToml != "success" {
        nadhi.Error("Uh oh! " + resultToml)
        return "Error: " + resultToml
    }

    nadhi.Success("Downloaded package " + packageArg + "..")
    return "Downloading package: " + packageArg
}