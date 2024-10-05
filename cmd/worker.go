package cmd

import (
	"temporal-scaffolding/pkg/config"
	"temporal-scaffolding/pkg/di"
	"temporal-scaffolding/pkg/worker"

	"github.com/spf13/cobra"
)

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "start worker",
	Long:  "Starts worker",
	Run: func(cmd *cobra.Command, args []string) {
		config := config.NewConfig()
		di, err := di.NewDI(config)
		if err != nil {
			panic(err)
		}

		worker := worker.NewWorker(di)
		if err := worker.Start(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(workerCmd)
}
