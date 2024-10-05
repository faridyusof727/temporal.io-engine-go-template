package cmd

import (
	"temporal-scaffolding/pkg/config"
	"temporal-scaffolding/pkg/di"
	"temporal-scaffolding/pkg/logger"
	"temporal-scaffolding/workflow/sample"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"go.temporal.io/sdk/client"
)

var temporalClientSampleCmd = &cobra.Command{
	Use:   "sample",
	Short: "sample",
	Long:  "sample",
	Run: func(cmd *cobra.Command, args []string) {
		config := config.NewConfig()
		di, err := di.NewDI(config)
		if err != nil {
			panic(err)
		}

		l, err := di.LoadLogger()
		if err != nil {
			panic(err)
		}

		temporalLogger := logger.NewTemporalLoggerAdapter(l)
		temporalClient, err := client.Dial(client.Options{
			HostPort: client.DefaultHostPort,
			Logger:   temporalLogger,
		})
		if err != nil {
			l.ErrorF("Unable to create Temporal Client", err)
		}
		defer temporalClient.Close()

		workflowOptions := client.StartWorkflowOptions{
			ID:        uuid.New().String(),
			TaskQueue: "default",
		}

		workflowRun, err := temporalClient.ExecuteWorkflow(cmd.Context(), workflowOptions, sample.SampleWorkflow)
		if err != nil {
			l.ErrorF("Unable to execute workflow", err)
		}

		var result string
		err = workflowRun.Get(cmd.Context(), &result)
		if err != nil {
			l.ErrorF("Unable to get workflow result", err)
		}

		l.InfoF("Workflow result: %s", result)
	},
}

func init() {
	rootCmd.AddCommand(temporalClientSampleCmd)
}
