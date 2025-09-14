package nadhi

import (
	"fmt"
	"os/exec"
	"strings"

	
)

func Install(apiURL, packageArg string) string {
	SuccessCheck("Installing package " + "..")
    packageName := packageArg
    version := ""
    if strings.Contains(packageArg, "@") {
        parts := strings.SplitN(packageArg, "@", 2)
        packageName = parts[0]
        version = parts[1]
    }

    pkgData, err := FetchPackages(apiURL)
	SuccessCheck("Refetching packages" + "..")

    if err != nil {
        return fmt.Sprintf("Error fetching package info: %v", err)
    }

    pkg, ok := pkgData.Packages[packageName]
    if !ok {
		
        return "Package not found: " + packageName
    }

    versions := pkg.(map[string]interface{})["versions"].(map[string]interface{})
    importPath := ""

	SuccessCheck("Checking a version for package" + "..")
    // Handle version logic
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

        return "Version not found for package: " + packageName
    }

    goGetArg := importPath
    if version != "" && version != "latest" {
        // If version is not just a major, append as @version
        goGetArg = fmt.Sprintf("%s@%s", importPath, version)
    }
	SuccessCheck(`Installing package via "go get"` + "..")
	cmd := exec.Command("go", "get", goGetArg)
	output, err := cmd.CombinedOutput()
	SuccessCheck("Running command: go get " + goGetArg)
	fmt.Printf("\033[90m%s\033[0m\n", string(output)) // Print output in grey
	if err != nil {
		return fmt.Sprintf("go get failed: %s", string(output))
	}

    return "Installed package: " + packageArg
}