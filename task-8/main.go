package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type CustomWaitGroup struct {
	sem   chan struct{}
	done  chan struct{}
	count int64
}

func NewCustomWG() *CustomWaitGroup {
	return &CustomWaitGroup{
		sem:  make(chan struct{}, 1),
		done: make(chan struct{}),
	}
}

func (wg *CustomWaitGroup) Add(delta int) {
	wg.sem <- struct{}{}
	if wg.count == 0 && delta > 0 {
		wg.done = make(chan struct{})
	}
	wg.count += int64(delta)
	if wg.count < 0 {
		<-wg.sem
		panic("CustomWaitGroup: negative counter")
	}
	if wg.count == 0 {
		close(wg.done)
	}
	<-wg.sem
}

func (wg *CustomWaitGroup) Done() {
	wg.Add(-1)
}

func (wg *CustomWaitGroup) Wait() {
	for {
		wg.sem <- struct{}{}
		done := wg.count == 0
		currentDone := wg.done
		<-wg.sem
		if done {
			return
		}
		<-currentDone
	}
}

func main() {
	cwg := NewCustomWG()
	var completed int64
	cwg.Add(10)

	for i := range 10 {
		go func(id int) {
			defer cwg.Done()
			time.Sleep(time.Duration(id) * 10 * time.Millisecond)
			atomic.AddInt64(&completed, 1)
		}(i)
	}

	cwg.Wait()
	fmt.Println("completed:", completed)
}
