package worker

import (
	sampleActivity "temporal-scaffolding/activity/sample"
	"temporal-scaffolding/pkg/di"
	"temporal-scaffolding/pkg/logger"
	sampleWorkflow "temporal-scaffolding/workflow/sample"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

type WorkerImpl struct {
	di *di.DI
}

func NewWorker(di *di.DI) *WorkerImpl {
	return &WorkerImpl{di: di}
}

func (w *WorkerImpl) Start() error {
	l, err := w.di.LoadLogger()
	if err != nil {
		return err
	}

	// We could inject the logger into the client options
	temporalClient, err := client.Dial(client.Options{
		Logger: logger.NewTemporalLoggerAdapter(l.(*logger.LoggerImpl)),
	})
	if err != nil {
		l.ErrorF("Unable to create client", err)
		return err
	}
	defer temporalClient.Close()

	// Create a new Worker.
	wk := worker.New(temporalClient, "default", worker.Options{})

	// Register your Workflow Definitions with the Worker.
	wk.RegisterWorkflow(sampleWorkflow.SampleWorkflow)

	// Register your Activity Definitions with the Worker.
	sampleActivity := &sampleActivity.SampleActivity{
		Parameter: "John Doe",
	}
	wk.RegisterActivity(sampleActivity)

	// Run the Worker
	err = wk.Run(worker.InterruptCh())
	if err != nil {
		l.ErrorF("Unable to start Worker", err)
		return err
	}

	return nil
}
