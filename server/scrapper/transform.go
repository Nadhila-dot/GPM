package scrapper

import (
    "strings"
)

// TransformedPackage is the minimal structure for client packages.json
type TransformedPackage struct {
    Versions struct {
        Latest   string            `json:"latest"`
        Default  string            `json:"default"`
        Major    map[string]string `json:"major"`
        Patch    []string          `json:"patch"`
    } `json:"versions"`
}

// TransformToClientFormat transforms parsed packages to the minimal client format.
func TransformToClientFormat(parsed []PackageResult) map[string]TransformedPackage {
    result := make(map[string]TransformedPackage)
    for _, pkg := range parsed {
        name := extractNameFromImport(pkg.ImportPath)
        tp := TransformedPackage{}
        tp.Versions.Latest = pkg.ImportPath
        tp.Versions.Default = pkg.ImportPath
        tp.Versions.Major = make(map[string]string)
        tp.Versions.Patch = []string{}
        // Try to extract major version from import path (e.g., /v2, /v3)
        major := extractMajorVersion(pkg.ImportPath)
        if major != "" {
            tp.Versions.Major[major] = pkg.ImportPath
        } else {
            tp.Versions.Major["1"] = pkg.ImportPath
        }
        result[name] = tp
    }
    return result
}

// extractNameFromImport gets the last path segment as the package name
func extractNameFromImport(importPath string) string {
    parts := strings.Split(importPath, "/")
    if len(parts) == 0 {
        return importPath
    }
    // If last part is a version (v2, v3), use previous
    last := parts[len(parts)-1]
    if strings.HasPrefix(last, "v") && len(parts) > 1 {
        return parts[len(parts)-2]
    }
    return last
}

// extractMajorVersion tries to get the major version from the import path (e.g., v2, v3)
func extractMajorVersion(importPath string) string {
    parts := strings.Split(importPath, "/")
    if len(parts) == 0 {
        return ""
    }
    last := parts[len(parts)-1]
    if strings.HasPrefix(last, "v") && len(last) > 1 {
        return strings.TrimPrefix(last, "v")
    }
    return ""
}
