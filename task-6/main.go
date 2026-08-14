package main

import (
	"context"
	"fmt"
	"math/rand"
)

func randomGenerator(ctx context.Context) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			value := rand.Int()
			select {
			case out <- value:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rgenCh := randomGenerator(ctx)
	for range 5 {
		fmt.Println(<-rgenCh)
	}
	
	cancel()
}
