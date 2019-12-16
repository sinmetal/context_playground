package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGoodbyeGoroutine(t *testing.T) {
	ctx := context.Background()

	tctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()

	err := GoodbyeGoroutine(tctx)
	if err != nil {
		fmt.Printf("err... %+v\n", err)
	}
}
