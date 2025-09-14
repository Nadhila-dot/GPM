package nadhi

func ExtractDefaultImport(packageInfo interface{}) string {
	infoMap, ok := packageInfo.(map[string]interface{})
	if !ok {
		return ""
	}

	versions, ok := infoMap["versions"].(map[string]interface{})
	if !ok {
		return ""
	}

	defaultVersion, ok := versions["default"].(string)
	if !ok {
		return ""
	}

	return defaultVersion
}