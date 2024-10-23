package worker

import (
	"temporal-scaffolding/activity/sample"
	"temporal-scaffolding/pkg/di"
)

func GetActivities(di *di.DI) map[string]interface{} {
	activities := map[string]interface{}{
		"SampleActivity": &sample.SampleActivity{
			Parameter: "John Doe",
		},
	}

	return activities
}
