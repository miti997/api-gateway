package bootstrap

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config interface {
	Load(configPath string) error
}

type ServerConfig struct {
	Address string `json:"address"`
}

type Route struct {
	Request string `json:"request"`
	In      string `json:"in"`
	Out     string `json:"out"`
}

type RoutesConfig struct {
	Routes []Route `json:"routes"`
}

type LoggerConfig struct {
	MaxSizeMB int    `json:"maxSizeMB"`
	FilePath  string `json:"filePath"`
	FileName  string `json:"fileName"`
}

func (sc *ServerConfig) Load(configPath string) error {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("could not read config file: %v", err)
	}

	err = json.Unmarshal(data, sc)
	if err != nil {
		return fmt.Errorf("could not parse server config file: %v", err)
	}
	return nil
}

func (rc *RoutesConfig) Load(configPath string) error {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("could not read routes config file: %v", err)
	}

	err = json.Unmarshal(data, rc)
	if err != nil {
		return fmt.Errorf("could not parse routes config file: %v", err)
	}
	return nil
}

func (lc *LoggerConfig) Load(configPath string) error {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("could not read logger config file: %v", err)
	}

	err = json.Unmarshal(data, lc)
	if err != nil {
		return fmt.Errorf("could not parse logger config file: %v", err)
	}
	return nil
}
