package main

import (
	"sync/atomic"
	"testing"
)

func TestCustomWgWaitsForGoroutines(t *testing.T) {
	wg := NewCustomWG()
	var completed int64
	wg.Add(5)

	for range 5 {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&completed, 1)
		}()
	}

	wg.Wait()
	if got := atomic.LoadInt64(&completed); got != 5 {
		t.Errorf("got %d completed goroutines, want 5", got)
	}
}

func TestCustomWgZero(t *testing.T) {
	wg := NewCustomWG()
	wg.Wait()
}
