package nadhi

import (
	"fmt"
	"os/exec"

	//"github.com/fatih/color"
	
)

// Uninstall removes a Go package using 'go get <importPath>@none'
// This is a sigma way to remove the go package cuz who wants to remove it manually in the go.mod and go.sum files

func Uninstall(importPath string) string {
	//fmt.Println(color.New(color.FgRed, color.Bold).Sprint("Uninstalling package " + importPath + ".."))
	RedLoading("Uninstalling package " + importPath + "..")
	cmd := exec.Command("go", "get", importPath+"@none")
	output, err := cmd.CombinedOutput()
	// ANSI escape code for grey: \033[90m ... \033[0m
	fmt.Printf("\033[90m%s\033[0m", string(output))
	if err != nil {
		Error(fmt.Sprintf("Error uninstalling package %s: %v", importPath, err))
		return "false"
	}
	return "true"
}