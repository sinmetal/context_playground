package main

import (
	"context"
	"fmt"
	"time"
)

func GoodbyeGoroutine(ctx context.Context) error {
	retCh := make(chan error)
	go func() {
		ret := fat()
		select {
		case <-ctx.Done():
		case retCh <- ret:
		}
	}()
	select {
	case <-ctx.Done():
		return ctx.Err()
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
