package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestStepByStep(t *testing.T) {
	ctx := context.Background()

	tctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()

	err := StepByStep(tctx)
	if err != nil {
		fmt.Printf("err... %+v\n", err)
	}
}
