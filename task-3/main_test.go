package main

import "testing"

func TestStringIntMap(t *testing.T) {

	m := NewSIMap()
	m.Add("test", 10)

	t.Run("test Add", func(t *testing.T) {
		if v, ok := m.Get("test"); !ok || v != 10 {
			t.Errorf("got (%d, %v), want (2, true)", v, ok)
		}

	})
	t.Run("test Get", func(t *testing.T) {
		if v, ok := m.Get("test2"); ok {
			t.Errorf("got (%d, %v), want (_, false)", v, ok)
		}
	})
	t.Run("test Exists", func(t *testing.T) {
		if ok := m.Exists("test"); !ok {
			t.Errorf("got %v, want true", ok)
		}
		if ok := m.Exists("test2"); ok {
			t.Errorf("got %v, want false", ok)
		}
	})
	m.Remove("test")
	t.Run("test Remove", func(t *testing.T) {
		if ok := m.Exists("test"); ok {
			t.Errorf("got %v, want false", ok)
		}
	})
	cpy := m.Copy()
	cpy.Add("test", 10)
	t.Run("test Copy", func(t *testing.T) {
		if ok := m.Exists("test"); ok {
			t.Errorf("got %v, want false", ok)
		}
	})
}
