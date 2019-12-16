package main

import (
	"context"
	"fmt"
	"time"
)

func StepByStep(ctx context.Context) error {
	fmt.Println("longRun sleep start")

	// contextを細かくチェックして、cancelされてたら、終了する
	for i := 0; i < 1000; i++ {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err(), "*child done*")
			return ctx.Err()
		default:
			fmt.Println(ctx.Err(), "child sleep")
			time.Sleep(10 * time.Millisecond)
		}
	}

	fmt.Println("longRun sleep end")
	return nil
}
