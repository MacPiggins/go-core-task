package main

import (
	"reflect"
	"testing"
)

func TestMakeFilter(t *testing.T) {
	t.Run("default test", func(t *testing.T) {
		got := makeFilter([]string{"test"})
		want := map[string]struct{}{
			"test": {},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("default test", func(t *testing.T) {
		got := makeFilter([]string{})
		want := make(map[string]struct{})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestFilter(t *testing.T) {
	t.Run("default test", func(t *testing.T) {
		got := filter(
			[]string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			[]string{"banana", "date", "fig"},
		)
		want := []string{"apple", "cherry", "43", "lead", "gno1"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("empty filter", func(t *testing.T) {
		got := filter(
			[]string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			[]string{},
		)
		want := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("full filter", func(t *testing.T) {
		got := filter(
			[]string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			[]string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
		)
		want := []string{}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}
