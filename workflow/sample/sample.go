package sample

import (
	"temporal-scaffolding/activity/sample"
	"time"

	"go.temporal.io/sdk/workflow"
)

func SampleWorkflow(ctx workflow.Context) (*string, error) {
	activityOptions := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}

	ctx = workflow.WithActivityOptions(ctx, activityOptions)

	var sampleActivity *sample.SampleActivity

	var activityResult string

	err := workflow.ExecuteActivity(ctx, sampleActivity.HelloWorld).Get(ctx, &activityResult)
	if err != nil {
		return nil, err
	}

	return &activityResult, nil
}
