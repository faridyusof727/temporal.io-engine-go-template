package worker

import (
	"temporal-scaffolding/workflow/sample"
)

var workflows = map[string]interface{}{
	"SampleWorkflow": sample.SampleWorkflow,
}

