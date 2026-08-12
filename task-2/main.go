package main

import (
	"fmt"
	"math/rand"
)

func randSlice(n int) []int {
	s := []int{}
	for i := 0; i < n; i++ {
		s = append(s, rand.Int())
	}
	return s
}

func sliceExample(s []int) []int {
	for i, v := range s {
		if v%2 != 0 {
			s = removeElement(s, i)
		}
	}
	return s
}

func addElements(s []int, v int) []int {
	s = append(s, v)
	return s
}

func copySlice(s []int) []int {
	newS := make([]int, len(s))
	copy(newS, s)
	return newS
}

func removeElement(s []int, i int) []int {
	if i < len(s)-1 {
		s = append(s[:i], s[i+1:]...)
	} else {
		s = s[:len(s)-1]
	}
	return s
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(s)
	s = randSlice(10)
	fmt.Println(s)
	s = sliceExample(s)
	fmt.Println(s)
	s = addElements(s, 1234)
	fmt.Println(s)
	newS := copySlice(s)
	s = removeElement(s, len(s)-1)
	fmt.Println("s: ",s)
	fmt.Println("newS: ",newS)
}
