package main

import "fmt"

func cubicPipe(input <-chan uint8, output chan<- float64) {
	defer close(output)
	for value := range input {
		converted := float64(value)
		output <- converted * converted * converted
	}
}

func main() {
	input := make(chan uint8)
	output := make(chan float64)

	go cubicPipe(input, output)

	go func() {
		defer close(input)
		for _, value := range []uint8{1, 2, 3, 4, 5} {
			input <- value
		}
	}()

	for value := range output {
		fmt.Println(value)
	}
}
