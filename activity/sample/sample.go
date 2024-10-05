package sample

import (
	"context"
	"fmt"
)

type SampleActivity struct {
	Parameter string
}

func (a *SampleActivity) HelloWorld(ctx context.Context) (string, error) {
	result := fmt.Sprintf("Hello World %s", a.Parameter)
	return result, nil
}
