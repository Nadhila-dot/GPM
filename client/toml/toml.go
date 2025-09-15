package toml

import (
	"bufio"
	"os"
	"strings"

	nadhi "nadhi.dev/binaries/gpm/helpers"
)

// Config represents the structure of the [Configuration] section in gpm.toml
type Config struct {
	Source   string
	Cache    map[string]string
	OS       string
	Metadata []string
}

// ParseConfigFromToml reads gpm.toml and parses the [Configuration] section.
// If required keys are missing, they are added with default values.
func ParseConfigFromToml() *Config {
    cfg := &Config{
        Source:   "https://go.pkg.lat",
        Cache:    make(map[string]string),
        OS:       "",
        Metadata: []string{},
    }
    file, err := os.Open("gpm.toml")
    if err != nil {
        // If file doesn't exist, create it with defaults
        osSys := nadhi.WhatOs()
        _ = os.WriteFile("gpm.toml", []byte("[Configuration]\nsource = \"https://go.pkg.lat\"\nOS = \""+osSys+"\"\nmetadata = \"\"\n\n[packages]\n"), 0644)
        return cfg
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    inConfig := false
    foundSource := false
    foundOS := false
    foundMetadata := false

    var lines []string
    for scanner.Scan() {
        line := scanner.Text()
        lines = append(lines, line)
        trimmed := strings.TrimSpace(line)
        if trimmed == "[Configuration]" {
            inConfig = true
            continue
        }
        if inConfig {
            if trimmed == "" || strings.HasPrefix(trimmed, "[") {
                inConfig = false
                continue
            }
            parts := strings.SplitN(trimmed, "=", 2)
            if len(parts) == 2 {
                key := strings.TrimSpace(parts[0])
                val := strings.Trim(strings.TrimSpace(parts[1]), "\"")
                switch key {
                case "source":
                    cfg.Source = val
                    foundSource = true
                case "OS":
                    cfg.OS = val
                    foundOS = true
                case "metadata":
                    cfg.Metadata = strings.Split(val, ",")
                    foundMetadata = true
                default:
                    cfg.Cache[key] = val
                }
            }
        }
    }

    // Add missing keys if needed
    changed := false
    if !foundSource || !foundOS || !foundMetadata {
        var newLines []string
        inConfig = false
        for _, line := range lines {
            newLines = append(newLines, line)
            if strings.TrimSpace(line) == "[Configuration]" {
                inConfig = true
                continue
            }
            if inConfig && (line == "" || strings.HasPrefix(strings.TrimSpace(line), "[")) {
                if !foundSource {
                    newLines = append(newLines, "source = \"https://go.pkg.lat\"")
                    changed = true
                }
                if !foundOS {
                    os := nadhi.WhatOs()
                    newLines = append(newLines, "OS = \""+os+"\"")
                    changed = true
                }
                if !foundMetadata {
                    newLines = append(newLines, "metadata = \"\"")
                    changed = true
                }
                inConfig = false
            }
        }
        // If config section is at the end of file
        if inConfig {
            if !foundSource {
                    newLines = append(newLines, "source = \"https://go.pkg.lat\"")
                    changed = true
                }
                if !foundOS {
                    os := nadhi.WhatOs()
                    newLines = append(newLines, "OS = \""+os+"\"")
                    changed = true
                }
                if !foundMetadata {
                    newLines = append(newLines, "metadata = \"\"")
                    changed = true
                }
        }
        if changed {
            _ = os.WriteFile("gpm.toml", []byte(strings.Join(newLines, "\n")), 0644)
        }
    }

    return cfg
}

// EditConfigValueInToml edits or adds a key-value pair in the [Configuration] section of gpm.toml
func EditConfigValueInToml(key, value string) error {
	const tomlFile = "gpm.toml"
	file, err := os.OpenFile(tomlFile, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	inConfig := false
	edited := false

	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)
		lines = append(lines, line)
		if trimmed == "[Configuration]" {
			inConfig = true
			continue
		}
		if inConfig {
			if trimmed == "" || strings.HasPrefix(trimmed, "[") {
				if !edited {
					lines = append(lines[:len(lines)-1], key+" = \""+value+"\"", line)
					edited = true
				}
				inConfig = false
				continue
			}
			if strings.HasPrefix(trimmed, key+" =") {
				lines[len(lines)-1] = key + " = \"" + value + "\""
				edited = true
			}
		}
	}
	if !edited && inConfig {
		lines = append(lines, key+" = \""+value+"\"")
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return os.WriteFile(tomlFile, []byte(strings.Join(lines, "\n")), 0644)
}

func CreateGpmFile() string {
	if _, err := os.Stat("gpm.toml"); err == nil {
		// gpm.toml file already exists
		return "success"
	}

	file, err := os.Create("gpm.toml")
	if err != nil {
		return "failed to create gpm.toml file: " + err.Error()
	}
	defer file.Close()

	_, err = file.WriteString(`# Go Package Manager TOML
# This file is used by gpm to manage packages
# Do not edit this file manually unless you know what you are doing
# visit go.pkg.lat to learn more

[packages]
`)
	if err != nil {
		return "failed to write to gpm.toml file: " + err.Error()
	}

	return "success"
}

func TomlAdd(name, link string) string {
	const tomlFile = "gpm.toml"
	entry := name + " = \"" + link + "\"\n"

	// Read the existing file
	file, err := os.OpenFile(tomlFile, os.O_RDWR, 0644)
	if err != nil {
		return "failed to open gpm.toml: " + err.Error()
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	inPackages := false
	inserted := false

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		if strings.TrimSpace(line) == "[packages]" {
			inPackages = true
			continue
		}
		if inPackages {
			trimmed := strings.TrimSpace(line)
			if strings.HasPrefix(trimmed, name+" =") {
				return "Package already exists.. Has it been installed already?"
			}
			// If we reach a new section or empty line after [packages], insert if not already inserted
			if !inserted && (trimmed == "" || strings.HasPrefix(trimmed, "[")) {
				lines = append(lines[:len(lines)-1], "[packages]", entry, line)
				inserted = true
				inPackages = false
			}
		}
	}
	if !inserted && inPackages {
		lines = append(lines, entry)
	}

	if err := scanner.Err(); err != nil {
		return "failed to read gpm.toml: " + err.Error()
	}

	// Write back to the file
	err = os.WriteFile(tomlFile, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		return "failed to write to gpm.toml: " + err.Error()
	}

	return "success"
}

func TomlRemove(name string) string {
	const tomlFile = "gpm.toml"

	// Read the existing file
	file, err := os.OpenFile(tomlFile, os.O_RDWR, 0644)
	if err != nil {
		return "failed to open gpm.toml: " + err.Error()
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	found := false

	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)
		// Skip the line if it matches the package entry
		if strings.HasPrefix(trimmed, name+" =") {
			found = true
			continue
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return "failed to read gpm.toml: " + err.Error()
	}

	if !found {
		return "Package not found"
	}

	// Write back to the file
	err = os.WriteFile(tomlFile, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		return "failed to write to gpm.toml: " + err.Error()
	}

	return "success"
}

func CheckToml() bool {
	if _, err := os.Stat("gpm.toml"); err == nil {
		return true // gpm.toml exists
	}
	return false // gpm.toml does not exist
}
func GetPackages() map[string]string {
	packages := make(map[string]string)
	file, err := os.Open("gpm.toml")
	if err != nil {
		return packages // return empty map if file can't be opened
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	inPackages := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "[packages]" {
			inPackages = true
			continue
		}
		if inPackages {
			if line == "" || strings.HasPrefix(line, "[") {
				break // end of packages section
			}
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				name := strings.TrimSpace(parts[0])
				link := strings.Trim(strings.TrimSpace(parts[1]), "\"")
				packages[name] = link
			}
		}
	}
	return packages
}

func GetImportPath(name string) string {
	file, err := os.Open("gpm.toml")
	if err != nil {
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	inPackages := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "[packages]" {
			inPackages = true
			continue
		}
		if inPackages {
			if line == "" || strings.HasPrefix(line, "[") {
				break // end of packages section
			}
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				pkgName := strings.TrimSpace(parts[0])
				importPath := strings.Trim(strings.TrimSpace(parts[1]), "\"")
				if pkgName == name {
					return importPath
				}
			}
		}
	}
	return ""
}
