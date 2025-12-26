package config

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

type Config struct {
	OnlySuggest  bool   `json:"only_suggest"`
	DefaultModel string `json:"default_model"`
	DebugMode    bool   `json:"debug_mode"`
}

type Retriever interface {
	Get() (error, Config)
}

type Commander interface {
	Set(cfg Config) error
}

type DefaultRetriever struct{}
type DefaultCommander struct{}

func NewRetriever() Retriever {
	return DefaultRetriever{}
}

func NewCommander() Commander {
	return DefaultCommander{}
}

func (retriever DefaultRetriever) Get() (error, Config) {
	configPath, err := Path()
	if err != nil {
		return err, Config{}
	}

	file, err := os.Open(configPath)
	if err != nil {
		return err, Config{}
	}

	defer closeFile(file)

	byteValue, err := io.ReadAll(file)

	if err != nil {
		return err, Config{}
	}

	var config Config
	err = json.Unmarshal(byteValue, &config)

	if err != nil {
		return err, Config{}
	}

	return nil, config
}

func (commander DefaultCommander) Set(cfg Config) error {
	configPath, err := Path()
	if err != nil {
		return err
	}

	file, err := os.OpenFile(configPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	defer closeFile(file)

	byteValue, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(byteValue)
	if err != nil {
		return err
	}

	return nil
}

func Path() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".whai", "config.json"), nil
}

func closeFile(f io.Closer) {
	if err := f.Close(); err != nil {
	}
}
