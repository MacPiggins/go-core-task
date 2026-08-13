package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestMerge(t *testing.T) {
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 3
	ch2 <- 2
	ch2 <- 4
	close(ch1)
	close(ch2)

	var got []int
	for value := range merge(ch1, ch2) {
		got = append(got, value)
	}

	want := []int{1, 2, 3, 4}
	sort.Ints(got)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestMergeEmpty(t *testing.T) {
	ch := make(chan int)
	close(ch)
	if _, ok := <-merge(ch); ok {
		t.Fatal("merged channel should be closed")
	}
}
