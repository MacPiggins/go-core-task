package main

import (
	"fmt"
	"sync"
)

func merge(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, ch := range channels {
		go func(input <-chan int) {
			defer wg.Done()
			for value := range input {
				out <- value
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer close(ch1)
		for _, value := range []int{1, 2, 3} {
			ch1 <- value
		}
	}()

	go func() {
		defer close(ch2)
		for _, value := range []int{4, 5, 6} {
			ch2 <- value
		}
	}()

	go func() {
		defer close(ch3)
		for _, value := range []int{7, 8, 9} {
			ch3 <- value
		}
	}()

	for value := range merge(ch1, ch2, ch3) {
		fmt.Println(value)
	}
}
