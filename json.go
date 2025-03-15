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

	// Decode the JSON file into a Config struct.
	if err := json.Unmarshal(data, &cfg); err != nil {
		return cfg, fmt.Errorf("unable to parse config file: %v", err)
	}
	// We re-write the config file to ensure all fields are present.
	if updated, err := WriteJSON(file, cfg); err != nil {
		return cfg, fmt.Errorf("unable to re-write config file: %v", err)
	} else if updated {
		// If the config file was updated, we need to re-parse it to ensure the Config struct is up-to-date.
		return ParseJSON(file)
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

	// Write default config to file.
	dat, err := json.MarshalIndent(DefaultConfig, "", "	")
	if err != nil {
		return fmt.Errorf("unable to write default config to file: %v", err)
	}

	if err := os.WriteFile(file, dat, 0644); err != nil {
		return fmt.Errorf("unable to write default config to file: %v", err)
	}
	return nil
}

// WriteJSON writes a Config struct to a JSON file.
func WriteJSON(file string, cfg Config) (bool, error) {
	var updated bool

	switch cfg.Version {
	case "": // The first version of the config did not have the version field.
		newCfg := DefaultConfig
		newCfg.AuthKey = cfg.AuthKey
		newCfg.Branch = cfg.Branch
		newCfg.LocalAddress = cfg.LocalAddress
		newCfg.RemoteAddress = cfg.RemoteAddress
		cfg = newCfg
		updated = true
	case "0.1-beta":
		newCfg := DefaultConfig
		newCfg.Movement.PersuasionThreshold = 0.001
		cfg = newCfg
		updated = true
	}

	dat, err := json.MarshalIndent(cfg, "", "	")
	if err != nil {
		return updated, fmt.Errorf("unable to write config to file: %v", err)
	}

	if err := os.WriteFile(file, dat, 0644); err != nil {
		return updated, fmt.Errorf("unable to write config to file: %v", err)
	}
	return updated, nil
}
