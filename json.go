package oconfig

import (
	"fmt"
	"os"

	"github.com/hjson/hjson-go/v4"
)

// ParseJSON parses a JSON file and returns a Config struct.
func ParseJSON(file string) error {
	var parsedCfg Config
	data, err := os.ReadFile(file)
	if err != nil {
		if err = CreateJSON(file); err != nil {
			return err
		}

		return fmt.Errorf("config file created - please fill in required fields")
	}

	// Decode the JSON file into a Config struct.
	if err := hjson.Unmarshal(data, &parsedCfg); err != nil {
		return fmt.Errorf("unable to parse config file: %v", err)
	}
	Cfg = parsedCfg
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
