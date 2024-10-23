package worker

import (
	"temporal-scaffolding/pkg/di"
	"temporal-scaffolding/pkg/logger"

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
		Logger: logger.NewTemporalLoggerAdapter(l),
	})
	if err != nil {
		l.ErrorF("Unable to create client", err)
		return err
	}
	defer temporalClient.Close()

	// Create a new Worker.
	wk := worker.New(temporalClient, "default", worker.Options{})

	// Register your Workflow Definitions with the Worker.
	for name, workflow := range workflows {
		l.InfoF("Registering workflow: %s", name)
		wk.RegisterWorkflow(workflow)
	}

	// Register your Activity Definitions with the Worker.
	for name, activity := range GetActivities(w.di) {
		l.InfoF("Registering activity: %s", name)
		wk.RegisterActivity(activity)
	}

	// Run the Worker
	err = wk.Run(worker.InterruptCh())
	if err != nil {
		l.ErrorF("Unable to start Worker", err)
		return err
	}

	return nil
}
