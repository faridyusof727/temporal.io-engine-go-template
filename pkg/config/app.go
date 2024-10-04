package config

type App struct {
	ProgramName string `mapstructure:"PROGRAM_NAME"`
	TeamName    string `mapstructure:"TEAM_NAME"`
	ENV         string `mapstructure:"ENV"`
}
