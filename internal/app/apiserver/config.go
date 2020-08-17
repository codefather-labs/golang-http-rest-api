package apiserver

type Config struct {
	Bindaddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
}

func NewConfig() *Config {
	return &Config{
		Bindaddr: ":8000",
		LogLevel: "debug",
	}
}
