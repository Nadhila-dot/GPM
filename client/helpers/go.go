package nadhi

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	_"github.com/fatih/color"
)

func Tidy() string {
	//fmt.Println(color.New(color.FgRed, color.Bold).Sprint("Uninstalling package " + importPath + ".."))
	LoadingCheck("Compiling and Syncing your packages..")
	cmd := exec.Command("go", "mod", "tidy")
	output, err := cmd.CombinedOutput()
	// ANSI escape code for grey: \033[90m ... \033[0m
	fmt.Printf("\033[90m%s\033[0m", string(output))
	if err != nil {
		Error(fmt.Sprintf("Error tidying packages: %v", err))
		return "false"
	}
	return "true"
}

// verifyPackages runs 'go mod verify' and removes any package with errors.
func VerifyPackagesAndRemoveOnError() string {
	fmt.Println("")
	LoadingCheck("Starting verification of installed packages..")
	LoadingCheck("Verifying packages..")
    cmd := exec.Command("go", "mod", "verify")
    var out bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &out
	LoadingCheck("Running 'go mod verify' to check package integrity..")
    err := cmd.Run()
    output := out.String()

	LoadingCheck("Analyzing verification results..")
    if err != nil || !strings.Contains(output, "all modules verified") {
        // Try to extract the problematic package from the output
        lines := strings.Split(output, "\n")
		LoadingCheck("Detected issues during verification..")
        for _, line := range lines {
            if strings.Contains(line, "checksum mismatch") {
                // Example line: "github.com/foo/bar@v1.2.3: checksum mismatch"
                parts := strings.Fields(line)
                if len(parts) > 0 {
                    pkg := strings.Split(parts[0], "@")[0]
                    Uninstall(pkg)
					LoadingCheck("Removing problematic package..")
                    Error("Removed problematic package: " + pkg)
                    return "false"
                }
            }
        }
        Error("Module verification failed")
        return "false"
    }

	println("")

    Success("All modules safe and verified")

    return "true"
}


func Govar(print bool) string {
    if print {
        SuccessCheck("Checking for Go version...")
    }
    

    cmd := exec.Command("go", "version")
    output, err := cmd.CombinedOutput()
    if err != nil {
        Error(fmt.Sprintf("Error getting Go version: %v", err))
        return "false"
    }

    // Example output: "go version go1.25.1 darwin/arm64"
    fields := strings.Fields(string(output))
    version := ""
    if len(fields) >= 3 {
        version = fields[2]
    } else {
        version = "unknown"
    }

    if print {
        SuccessCheck(fmt.Sprintf("Go version: %s", version))
    }

    return version
}