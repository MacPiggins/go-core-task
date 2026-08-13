package main

import (
	"reflect"
	"testing"
)

func TestTtos(t *testing.T) {
	t.Run("test int", func(t *testing.T) {
		got := ttos(10)
		want := "int"
		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})
	t.Run("test float", func(t *testing.T) {
		got := ttos(10.10)
		want := "float64"
		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})
	t.Run("test string", func(t *testing.T) {
		got := ttos("test")
		want := "string"
		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})
	t.Run("test bool", func(t *testing.T) {
		got := ttos(true)
		want := "bool"
		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})
	t.Run("test int", func(t *testing.T) {
		got := ttos(2 + 3i)
		want := "complex128"
		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})
}

func TestVtoa(t *testing.T) {
	t.Run("test int", func(t *testing.T) {
		got := vtoa(10)
		want := "10"
		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})
	t.Run("test float", func(t *testing.T) {
		got := vtoa(10.10)
		want := "10.1"
		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})
	t.Run("test string", func(t *testing.T) {
		got := vtoa("test")
		want := "test"
		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})
	t.Run("test bool", func(t *testing.T) {
		got := vtoa(true)
		want := "true"
		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})
	t.Run("test int", func(t *testing.T) {
		got := vtoa(2 + 3i)
		want := "(2+3i)"
		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})
}

func TestCombine(t *testing.T) {
	t.Run("default test", func(t *testing.T) {
		got := combine(10, 10.1, "test", true, 2+3i)
		want := "1010.1testtrue(2+3i)"
		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})

}

func TestToRunes(t *testing.T) {
	t.Run("default test", func(t *testing.T) {
		got := toRunes("test")
		want := []rune("test")
		if reflect.DeepEqual(got, want) != true {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

}

func TestAddSalt(t *testing.T) {
	t.Run("default test", func(t *testing.T) {
		got := string(addSalt([]rune("testtest"), "salt"))
		want := "testsalttest"
		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})
}
