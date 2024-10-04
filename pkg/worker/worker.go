package worker

import (
	"temporal-scaffolding/pkg/config"
	"temporal-scaffolding/pkg/di"
	sampleWorkflow "temporal-scaffolding/workflow/sample"
	sampleActivity "temporal-scaffolding/activity/sample"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

type WorkerImpl struct {
	di *di.DI
	config *config.Config
}

func NewWorker(di *di.DI, config *config.Config) *WorkerImpl {
	return &WorkerImpl{di: di, config: config}
}

func (w *WorkerImpl) Start() error {
	// We could inject the logger into the client options
	temporalClient, err := client.Dial(client.Options{})
	if err != nil {
		w.di.Logger.ErrorF("Unable to create client", err)
		return err
	}
	defer temporalClient.Close()

	// Create a new Worker.
	wk := worker.New(temporalClient, "default", worker.Options{})
	
	// Register your Workflow Definitions with the Worker.
	wk.RegisterWorkflow(sampleWorkflow.SampleWorkflow)
	
	// Register your Activity Definitions with the Worker.
	sampleActivity := &sampleActivity.SampleActivity{}
	wk.RegisterActivity(sampleActivity)

	// Run the Worker
	err = wk.Run(worker.InterruptCh())
	if err != nil {
		w.di.Logger.ErrorF("Unable to start Worker", err)
		return err
	}

	return nil
}
