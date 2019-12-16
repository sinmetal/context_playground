package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func GoodbyeGoroutine(ctx context.Context) error {
	retCh := make(chan error)
	go func() {
		retCh <- fat()
	}()
	select {
	case <-ctx.Done():
		return errors.New("timeout")
	case ret := <-retCh:
		return ret
	}
}

func fat() error {
	fmt.Println("fat start")
	time.Sleep(10 * time.Second)
	fmt.Println("fat end")
	return nil
}
