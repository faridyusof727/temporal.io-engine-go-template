package di

import (
	"fmt"
	"temporal-scaffolding/pkg/config"
	"temporal-scaffolding/pkg/logger"
)

type DI struct {
	Logger logger.Logger
}

func NewDI(config *config.Config) (*DI, error) {
	// load logger
	logger, err := loadLogger(config)
	if err != nil {
		return nil, fmt.Errorf("failed to load logger: %w", err)
	}

	return &DI{
		Logger: logger,
	}, nil
}

func loadLogger(config *config.Config) (logger.Logger, error) {
	logger, err := logger.New(logger.Options{
		DefaultFields: &logger.DefaultFields{
			Program: config.App.ProgramName,
			Team:    config.App.TeamName,
			ENV:     config.App.ENV,
		},
		LogPath: config.Logger.LogPath,
	})
	if err != nil {
		return nil, err
	}

	return logger, nil
}
