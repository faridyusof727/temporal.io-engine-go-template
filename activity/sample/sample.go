package sample

import "context"

type SampleActivity struct {
	
}

func (a *SampleActivity) HelloWorld(ctx context.Context) (*string, error) {
	result := "Hello World"
	return &result, nil
}
