package main

import "fmt"

func makeSet(s []int) map[int]bool {
	set := make(map[int]bool)
	for _, v := range s {
		set[v] = true
	}
	return set
}

func intersection(s1, s2 []int) (b bool, res []int) {
	var set map[int]bool
	var s []int
	if len(s1) < len(s2) {
		set = makeSet(s1)
		s = s2
	} else {
		set = makeSet(s2)
		s = s1
	}

	b = false
	res = make([]int, 0)
	for _, v := range s {
		if _, ok := set[v]; ok {
			res = append(res, v)
			b = true
		}
	}
	return
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	fmt.Println(intersection(b, a))
}
