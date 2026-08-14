package main

import (
	"reflect"
	"testing"
)

func TestMakeSet(t *testing.T) {
	got := makeSet([]int{2, 3})
	want := map[int]bool{
		2: false,
		3: false,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

}

func TestIntersection(t *testing.T) {
	t.Run("default test", func(t *testing.T) {
		ok, got := intersection([]int{65, 3, 58, 678, 64}, []int{64, 2, 3, 43})
		if !ok {
			t.Error("expected an intersection")
		}
		want := []int{3, 64}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("no intersection test", func(t *testing.T) {
		ok, got := intersection([]int{1, 2}, []int{3, 4})
		if ok {
			t.Error("expected no intersection")
		}

		if len(got) != 0 {
			t.Errorf("got %v, want empty", got)
		}
	})

	t.Run("unique test", func(t *testing.T) {
		_, got := intersection([]int{1, 2, 2, 3, 2}, []int{2, 3})
		want := []int{2, 3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}
