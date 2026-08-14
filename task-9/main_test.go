package main

import (
	"reflect"
	"testing"
)

func TestCubicPipe(t *testing.T) {
	input := make(chan uint8, 4)
	output := make(chan float64)
	input <- 2
	input <- 3
	input <- 4
	input <- 5
	close(input)

	go cubicPipe(input, output)

	var got []float64
	for value := range output {
		got = append(got, value)
	}
	want := []float64{8, 27, 64, 125}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}