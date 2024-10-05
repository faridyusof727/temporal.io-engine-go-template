package config

type Logger struct {
	Level   int    `mapstructure:"LOG_LEVEL"`
	LogPath string `mapstructure:"LOG_PATH"`
}
