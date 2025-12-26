package config

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
)

type Config struct {
	OnlySuggest         bool                          `json:"only_suggest"`
	Model               string                        `json:"model"`
	DebugMode           bool                          `json:"debug_mode"`
	ModelsConfiguration map[string]ModelConfiguration `json:"models_configuration"`
}

type ModelConfiguration struct {
	URL    string `json:"url"`
	ApiKey string `json:"api_key"`
}

func (cfg Config) NewConfig() Config {
	return Config{
		OnlySuggest: false,
		DebugMode:   false,
		Model:       "openai",
		ModelsConfiguration: map[string]ModelConfiguration{
			"openai": {},
		},
	}
}

func (cfg Config) GetModelConfiguration(model string) (error, ModelConfiguration) {
	modelConfig := cfg.ModelsConfiguration[model]
	if modelConfig.URL == "" && modelConfig.ApiKey == "" {
		return errors.New("no model configuration found"), ModelConfiguration{}
	}
	return nil, modelConfig
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
	if err != nil {
		return err
	}
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
