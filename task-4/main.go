package main

import "fmt"

func makeFilter(s []string) map[string]struct{} {
	filter := make(map[string]struct{})
	for _, v := range s {
		filter[v] = struct{}{}
	}
	return filter
}

func filter(s1, s2 []string) []string {
	filter := makeFilter(s2)
	for i, v := range s1 {
		if _, ok := filter[v]; ok {
			s1 = append(s1[:i], s1[i+1:]...)
		}
	}
	return s1
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	fmt.Println(filter(slice1, slice2))
}
