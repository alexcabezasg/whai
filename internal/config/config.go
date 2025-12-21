package config

type Config struct {
	OnlySuggest  bool   `json:"only_suggest"`
	DefaultModel string `json:"default_model"`
}

type Provider interface {
	Get() (error, Config)
	Set(cfg Config) error
}

type DefaultProvider struct{}

func (provider DefaultProvider) Get() (error, Config) {
	return nil, Config{
		OnlySuggest:  false,
		DefaultModel: "",
	}
}

func (provider DefaultProvider) Set(cfg Config) error {
	return nil
}

func NewProvider() Provider {
	return DefaultProvider{}
}
