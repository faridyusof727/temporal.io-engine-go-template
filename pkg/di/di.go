package di

import (
	"temporal-scaffolding/pkg/config"
	"temporal-scaffolding/pkg/logger"
)

type DI struct {
	config *config.Config
}

func NewDI(config *config.Config) (*DI, error) {
	return &DI{
		config: config,
	}, nil
}

func (d *DI) LoadConfig() (*config.Config, error) {
	return d.config, nil
}

func (d *DI) LoadLogger() (logger.Logger, error) {
	logger, err := logger.New(logger.Options{
		DefaultFields: &logger.DefaultFields{
			Program: d.config.App.ProgramName,
			Team:    d.config.App.TeamName,
			ENV:     d.config.App.ENV,
		},
		LogPath: d.config.Logger.LogPath,
		Level:   d.config.Logger.Level,
	})
	if err != nil {
		return nil, err
	}

	return logger, nil
}
