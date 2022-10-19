package conf

type Config struct {
	Address string `toml:"address"`
}

func defaultConfig() *Config {
	return &Config{
		Address: ":8080",
	}
}

func Init() (*Config, error) {
	cfg := &Config{}
	cfg = defaultConfig()
	return cfg, nil
}
