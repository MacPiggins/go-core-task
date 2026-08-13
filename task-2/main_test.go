package main

import (
	"reflect"
	"testing"
)

func TestSliceExample(t *testing.T) {
	t.Run("default test", func(t *testing.T) {
		got := sliceExample([]int{1, 2, 3, 4, 6})
		want := []int{2, 4, 6}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("erase first", func(t *testing.T) {
		got := sliceExample([]int{1, 2, 2, 4, 6})
		want := []int{2, 2, 4, 6}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("erase last", func(t *testing.T) {
		got := sliceExample([]int{2, 2, 2, 4, 5})
		want := []int{2, 2, 2, 4}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("erase middle", func(t *testing.T) {
		got := sliceExample([]int{2, 1, 1, 1, 2})
		want := []int{2, 2}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("erase all", func(t *testing.T) {
		got := sliceExample([]int{1, 1, 1, 1, 1})
		want := []int{}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("erase empty", func(t *testing.T) {
		got := sliceExample([]int{})
		want := []int{}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestAddElements(t *testing.T) {
	t.Run("default test", func(t *testing.T) {
		input := []int{1, 2}
		got := addElements(input, 3)
		want := []int{1, 2, 3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("add to empty", func(t *testing.T) {
		input := []int{}
		got := addElements(input, 3)
		want := []int{3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestCopySliceIndependent(t *testing.T) {
	input := []int{1, 2, 3}
	copied := copySlice(input)
	copied[0] = 99
	if input[0] != 1 {
		t.Errorf("original slice changed: %v", input)
	}
}

func TestRemoveElement(t *testing.T) {
	t.Run("default test", func(t *testing.T) {
		got := removeElement([]int{10, 20, 30, 40}, 2)
		want := []int{10, 20, 40}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("remove first", func(t *testing.T) {
		got := removeElement([]int{10, 20, 30, 40}, 0)
		want := []int{20, 30, 40}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("remove last", func(t *testing.T) {
		got := removeElement([]int{10, 20, 30, 40}, 3)
		want := []int{10, 20, 30}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("remove from empty", func(t *testing.T) {
		got := removeElement([]int{}, 2)
		want := []int{}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("invalid index", func(t *testing.T) {
		got := removeElement([]int{10, 20, 30, 40}, 10)
		want := []int{10, 20, 30, 40}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}
