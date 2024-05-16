package oconfig

import (
	"encoding/json"
	"fmt"
	"os"
)

// ParseJSON parses a JSON file and returns a Config struct.
func ParseJSON(file string) (Config, error) {
	var cfg Config
	data, err := os.ReadFile(file)
	if err != nil {
		if err = CreateJSON(file); err != nil {
			return cfg, err
		}

		return DefaultConfig, fmt.Errorf("config file created - please fill in required fields")
	}

	var configData map[string]interface{}
	if err := json.Unmarshal(data, &configData); err != nil {
		return cfg, fmt.Errorf("unable to parse config file: %v", err)
	}

	if err := ValidateJSON(&cfg, configData); err != nil {
		return cfg, err
	}

	return cfg, nil
}

// CreateJSON creates a new JSON file with default config.
func CreateJSON(file string) error {
	// Create a new config file
	_, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("unable to create config file: %v", err)
	}

	// Write default config to file
	dat, err := json.MarshalIndent(map[string]interface{}{
		"auth_key":    string(DefaultConfig.AuthKey[:]),
		"branch":      DefaultConfig.Branch,
		"local_addr":  DefaultConfig.LocalAddress,
		"remote_addr": DefaultConfig.RemoteAddress,
	}, "", "	")

	if err != nil {
		return fmt.Errorf("unable to write default config to file: %v", err)
	}

	if err := os.WriteFile(file, dat, 0644); err != nil {
		return fmt.Errorf("unable to write default config to file: %v", err)
	}
	return nil
}

// ValidateConfig validates the config data
func ValidateJSON(cfg *Config, data map[string]interface{}) error {
	if authKey, ok := data["auth_key"]; ok {
		copy(cfg.AuthKey[:], []byte(authKey.(string)))
	} else {
		return fmt.Errorf("auth_key field is missing from the config")
	}

	if branch, ok := data["branch"]; ok {
		cfg.Branch = branch.(string)
	} else {
		return fmt.Errorf("branch field is missing from the config")
	}

	if localAddress, ok := data["local_addr"]; ok {
		cfg.LocalAddress = localAddress.(string)
	} else {
		return fmt.Errorf("local_addr field is missing from the config")
	}

	if remoteAddress, ok := data["remote_addr"]; ok {
		cfg.RemoteAddress = remoteAddress.(string)
	} else {
		return fmt.Errorf("remote_addr field is missing from the config")
	}

	return nil
}
