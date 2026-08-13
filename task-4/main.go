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
	for i := 0; i < len(s1); i++ {
		if _, ok := filter[s1[i]]; ok {
			s1 = append(s1[:i], s1[i+1:]...)
			i--
		}
	}
	return s1
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	fmt.Println(filter(slice1, slice2))
}
