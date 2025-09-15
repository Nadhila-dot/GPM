package nadhi

import (
    "fmt"
    "os"
    "os/exec"
    "strings"
)

// Build builds the Go project for multiple OS/arch, skipping any in ignoreos.
func Build(appname string, dir string, ignoreos []string) {
	LoadingCheck("Building the project...")
	LoadingCheck(fmt.Sprintf("Application Name: %s\nLocation: %s", appname, dir))

	// Supported targets
	targets := []struct {
		OS   string
		Arch string
		Ext  string
	}{
		{"linux", "amd64", ""},
		{"darwin", "amd64", ""},
		{"darwin", "arm64", ""},
		{"windows", "amd64", ".exe"},
	}

	// Create build directory
	buildDir := "builds"
	if err := os.MkdirAll(buildDir, 0755); err != nil {
		Error(fmt.Sprintf("Failed to create build directory: %v", err))
		return
	}

	// Helper to check if OS is in ignoreos
	shouldIgnore := func(os string) bool {
		for _, ignore := range ignoreos {
			if strings.EqualFold(ignore, os) {
				return true
			}
		}
		return false
	}

	for _, target := range targets {
		if shouldIgnore(target.OS) {
			LoadingCheck(fmt.Sprintf("Skipping build for %s (%s) as per ignore list.", target.OS, target.Arch))
			continue
		}
		outFile := fmt.Sprintf("%s/%s-%s-%s%s", buildDir, appname, target.OS, target.Arch, target.Ext)
		LoadingCheck(fmt.Sprintf("Building for %s (%s)...", target.OS, target.Arch))
		cmd := exec.Command("go", "build", "-o", outFile, dir)
		cmd.Env = append(os.Environ(),
			"GOOS="+target.OS,
			"GOARCH="+target.Arch,
		)
		if err := cmd.Run(); err != nil {
			Error(fmt.Sprintf("Build failed for %s/%s: %v", target.OS, target.Arch, err))
		} else {
			SuccessCheck(fmt.Sprintf("Built: %s", outFile))
		}
	}
	Success(fmt.Sprintf("All builds complete. Binaries are in the '%s' folder.", buildDir))
}