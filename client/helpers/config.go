package nadhi

import (
    "encoding/json"
   
    "os"
    "path/filepath"
    "runtime"
)

// Config represents the structure of config.json
type Config struct {
    Source   string                 `json:"source"`
    Cache    map[string]string      `json:"cache"`
    OS       string                 `json:"OS"`
    Metadata []interface{}          `json:"metadata"`
}



// DefaultConfig returns a Config with default values
func DefaultConfig() *Config {
    exePath, _ := os.Executable()
    cacheDir := filepath.Join(filepath.Dir(exePath), "cache")
    return &Config{
        Source: "https://go.pkg.lat",
        Cache:  map[string]string{"directory": cacheDir},
        OS:     runtime.GOOS,
        Metadata: []interface{}{},
    }
}


func CreateConfigFile(path string) error {
    if _, err := os.Stat(path); err == nil {
        return nil // File already exists
    }
    cfg := DefaultConfig()
    data, err := json.MarshalIndent(cfg, "", "    ")
    if err != nil {
        return err
    }
    return os.WriteFile(path, data, 0644)
}


func EditConfigValue(path, key, value string) error {
	cfg, err := GetConfig(path)
	if err != nil {
		return err
	}
	switch key {
	case "source":
		cfg.Source = value
	case "OS":
		cfg.OS = value
	default:
		// Try cache subkey
		if cfg.Cache == nil {
			cfg.Cache = map[string]string{}
		}
		cfg.Cache[key] = value
	}
	data, err := json.MarshalIndent(cfg, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// GetConfig loads config.json from the current working directory and returns the Config struct
func GetConfig(_ string) (*Config, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	configPath := filepath.Join(cwd, "config.json")
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// GetConfigValue gets a value from config.json (top-level or cache subkey) in the current working directory
func GetConfigValue(key string) string {
	cfg, err := GetConfig("")
	if err != nil {
		return ""
	}
	switch key {
	case "source":
		return cfg.Source
	case "OS":
		return cfg.OS
	default:
		if v, ok := cfg.Cache[key]; ok {
			return v
		}
	}
	return ""
}