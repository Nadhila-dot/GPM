package commands

import (
	"fmt"

	"github.com/fatih/color"
	nadhi "nadhi.dev/binaries/gpm/helpers"
	"nadhi.dev/binaries/gpm/toml"
)

// We all hate formating...
// We should use gpt-4 for this
// Which is exactly what I did


func Refresh(searchQuery ...string) string {
	nadhi.Logo()
	config := toml.ParseConfigFromToml()
	if config == nil {
		color.Red("Error getting config")
		return ""
	}

	nadhi.LoadingCheck(fmt.Sprintf("Refreshing package list from %s...", config.Source))

	packages, err := fetchPackages(config.Source, searchQuery...)
	if err != nil {
		handleFetchError(err)
		return ""
	}

	displayPackages(packages, searchQuery...)
	return "Refreshing packages"
}

func fetchPackages(sourceURL string, searchQuery ...string) (*nadhi.PackageData, error) {
	if len(searchQuery) > 0 && searchQuery[0] != "" {
		return nadhi.SearchPackages(sourceURL, searchQuery[0])
	}
	return nadhi.FetchPackages(sourceURL)
}

func handleFetchError(err error) {
	errorMessage := err.Error()

	if nadhi.CheckString([]string{"invalid character"}, errorMessage) {
		nadhi.Error("Response from server was wrong, Are you sure this is a GPM server?")
	} else {
		color.Red("Error fetching packages: %v", err)
	}
}

func displayPackages(pkgData *nadhi.PackageData, searchQuery ...string) {
	if len(searchQuery) > 0 && searchQuery[0] != "" {
		color.Cyan("Listing available packages matching: %s", searchQuery[0])
	} else {
		color.New(color.FgCyan, color.Bold).Println("Listing all available packages:")
	}

	if len(pkgData.Packages) == 0 {
		color.Yellow("No packages found.")
		return
	}

	greenText := color.New(color.FgGreen).SprintFunc()
	blueText := color.New(color.FgBlue).SprintFunc()
	magentaText := color.New(color.FgMagenta).SprintFunc()
	yellowText := color.New(color.FgYellow).SprintFunc()

	for packageName, packageInfo := range pkgData.Packages {
		defaultImport := nadhi.ExtractDefaultImport(packageInfo)
		fmt.Printf("%s: %s\n", greenText(packageName), blueText(defaultImport))

		// Show version info if available
		if m, ok := packageInfo.(map[string]interface{}); ok {
			if versions, vok := m["versions"].(map[string]interface{}); vok {
				// Major versions
				if major, mok := versions["major"].(map[string]interface{}); mok && len(major) > 0 {
					majors := []string{}
					for k := range major {
						majors = append(majors, k)
					}
					fmt.Printf("  %s %s %s\n", magentaText("↳ Major:"), yellowText(fmt.Sprintf("%v", majors)), "")
				}
				// Latest
				if latest, lok := versions["latest"].(string); lok && latest != "" {
					fmt.Printf("  %s %s\n", magentaText("↳ Latest:"), yellowText(latest))
				}
				// Default
				if def, dok := versions["default"].(string); dok && def != "" {
					fmt.Printf("  %s %s\n", magentaText("↳ Default:"), yellowText(def))
				}
			}
		}
		fmt.Println()
	}
}
