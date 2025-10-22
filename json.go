package oconfig

import (
	"errors"
	"fmt"
	"os"

	"github.com/hjson/hjson-go/v4"
)

var (
	ErrConfigCreated = errors.New("config file created - please fill in required fields")
	ErrConfigUpdated = errors.New("config file updated - please fill in required fields")
)

// ParseRawJSON parses a raw JSON string and returns a Config struct.
func ParseRawJSON(data []byte) (Config, error) {
	parsedCfg := DefaultConfig
	if err := hjson.Unmarshal(data, &parsedCfg); err != nil {
		return Config{}, fmt.Errorf("unable to parse config: %w", err)
	}

	if parsedCfg.Version != ConfigVersion {
		newCfg := parsedCfg
		switch parsedCfg.Version {
		case 0: // No version set.
			newCfg.Prefix = DefaultConfig.Prefix
			newCfg.GCPercent = DefaultConfig.GCPercent
			newCfg.MemThreshold = DefaultConfig.MemThreshold
			newCfg.Detections = DefaultConfig.Detections
		case 2:
			newCfg.Network = DefaultConfig.Network
		case 3:
			// For version 4, we added two new detections to the configuration.
			newCfg.Detections["Proxy_A"] = DefaultConfig.Detections["Proxy_A"]
			newCfg.Detections["Proxy_B"] = DefaultConfig.Detections["Proxy_B"]
		}
		newCfg.Version = ConfigVersion
		return newCfg, ErrConfigUpdated
	}

	return parsedCfg, nil
}

// ParseJSON parses a JSON file and returns a Config struct.
func ParseJSON(file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		if err = CreateJSON(file); err != nil {
			return err
		}

		return ErrConfigCreated
	}

	parsedCfg, err := ParseRawJSON(data)
	if err != nil && !errors.Is(err, ErrConfigUpdated) {
		return fmt.Errorf("unable to parse config file: %w", err)
	}

	if errors.Is(err, ErrConfigUpdated) {
		if writeErr := WriteJSON(file, parsedCfg); writeErr != nil {
			return fmt.Errorf("unable to update config file: %w", writeErr)
		}
		return ErrConfigUpdated
	}

	if writeErr := WriteJSON(file, parsedCfg); writeErr != nil {
		return fmt.Errorf("unable to re-write config file: %w", writeErr)
	}

	Global = parsedCfg
	return nil
}

// CreateJSON creates a new JSON file with default config.
func CreateJSON(file string) error {
	// Create a new config file
	_, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("unable to create config file: %v", err)
	}

	// Write default config to file.
	dat, err := hjson.MarshalWithOptions(DefaultConfig, hjson.EncoderOptions{
		IndentBy:              "    ",
		EmitRootBraces:        true,
		QuoteAlways:           false,
		QuoteAmbiguousStrings: false,
		Eol:                   "\n",
		Comments:              true,
	})
	if err != nil {
		return fmt.Errorf("unable to write default config to file: %v", err)
	}

	if err := os.WriteFile(file, dat, 0644); err != nil {
		return fmt.Errorf("unable to write default config to file: %v", err)
	}
	return nil
}

func WriteJSON(file string, cfg Config) error {
	dat, err := hjson.MarshalWithOptions(cfg, hjson.EncoderOptions{
		IndentBy:              "    ",
		EmitRootBraces:        true,
		QuoteAlways:           false,
		QuoteAmbiguousStrings: false,
		Eol:                   "\n",
		Comments:              true,
	})
	if err != nil {
		return fmt.Errorf("unable to write config to file: %v", err)
	}

	if err := os.WriteFile(file, dat, 0644); err != nil {
		return fmt.Errorf("unable to write config to file: %v", err)
	}
	return nil
}
