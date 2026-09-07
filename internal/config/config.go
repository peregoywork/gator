package config

import (
	"os"
	"fmt"
	"encoding/json"
)

const configFilename = ".gatorconfig.json"

type Config struct {
	DbUrl 			string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

// Read contents of .gatorconfig.json
func Read() (Config, error) {
	filepath, err := getConfigFilePath()
	if err != nil {
		fmt.Println("error getting config filepath")
		return Config{}, err
	}

	content, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("error reading config file")
		return Config{}, err
	}

	var cfg Config
	err = json.Unmarshal(content, &cfg)
	if err != nil {
		fmt.Println("error unmarshalling json")
		return Config{}, err
	}

	return cfg, nil
}

// Set the current user, then writes to the .gatorconfig.json file 
func (cfg Config) SetUser(user string) {
	cfg.CurrentUserName = user
	err := write(cfg)
	if err != nil {
		fmt.Println("error write current user to config")
	}
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("error trying to access user home dir")
		return "", err
	}

	return homeDir + "/" + configFilename, nil
}

func write(cfg Config) error {
	filepath, err := getConfigFilePath()
	if err != nil {
		fmt.Println("error reading config file", err)
		return err
	}

	// write Config to filepath
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		fmt.Println("error marshalling config data")
		return err
	}

	err = os.WriteFile(filepath, data, 0644)
	if err != nil {
		fmt.Println("error writing file:", err)
		return err
	}

	return nil
}


