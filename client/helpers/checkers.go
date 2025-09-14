package nadhi

import (
	"os"
	"runtime"
)

// HasGoMod checks if a go.mod file exists in the current directory.
// Now the thing won't install unless a go.mod file is present.
func HasGoMod() bool {
    _, err := os.Stat("go.mod")
    return err == nil
}

func WhatOs() string {
	if runtime.GOOS == "darwin" {
		return "MacOS"
	} else if runtime.GOOS == "windows" {
		return "Windows"
	} else if runtime.GOOS == "linux" {
		return "Linux"
	}
	return runtime.GOOS
}