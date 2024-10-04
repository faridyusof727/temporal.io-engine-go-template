package cmd

import (
	"temporal-scaffolding/pkg/config"
	"temporal-scaffolding/pkg/di"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "Some example root command",
	Run: func(cmd *cobra.Command, args []string) {
		config := config.NewConfig()
		di, err := di.NewDI(config)
		if err != nil {
			panic(err)
		}

		di.Logger.InfoF("Welcome to Temporal Scaffolding")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
