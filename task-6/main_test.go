package main

import (
	"context"
	"testing"
)

func TestRandomGeneratorProducesValues(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := randomGenerator(ctx)
	for range 3 {
		select {
		case <-ch:
		case <-ctx.Done():
			t.Error("generator stopped too early")
		}
	}
}

func TestRandomGeneratorStopsOnCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ch := randomGenerator(ctx)
	cancel()

	select {
	case _, ok := <-ch:
		if ok {
			for range ch {
			}
		}
	case <-ctx.Done():
	}
}
