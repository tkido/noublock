package main

import (
	"testing"
)

func TestSetEqual(t *testing.T) {
	cases := []struct {
		want   bool
		s1, s2 Set
	}{
		{true, NewSet(), NewSet()},
		{false, NewSet(), NewSet(Vector{0, 0})},
		{true, NewSet(Vector{0, 0}, Vector{1, 1}), NewSet(Vector{0, 0}, Vector{1, 1})},
		{false, NewSet(Vector{1, 0}, Vector{1, 1}), NewSet(Vector{0, 0}, Vector{1, 1})},
		{true, NewSet(Vector{0, 0}, Vector{1, 1}), NewSet(Vector{1, 1}, Vector{0, 0})},
		{false, NewSet(Vector{0, 0}, Vector{1, 1}), NewSet(Vector{0, 0})},
		{true, NewSet(Vector{1, 2}, Vector{2, 4}), NewSet(Vector{1, 2}, Vector{2, 4})},
	}
	for _, c := range cases {
		got := c.s1.Equal(c.s2)
		if got != c.want {
			t.Errorf("got %v want %v", got, c.want)
		}
	}
}

func TestSetOrigin(t *testing.T) {
	cases := []struct {
		want  Vector
		given Set
	}{
		{Vector{1, -1}, NewSet(Vector{-1, 1}, Vector{1, -1})},
		{Vector{-1, -1}, NewSet(Vector{-1, -1}, Vector{3, 4})},
		{Vector{-1, -1}, NewSet(Vector{-1, -1}, Vector{0, -1})},
	}
	for _, c := range cases {
		got := c.given.Origin()
		if got != c.want {
			t.Errorf("got %v want %v", got, c.want)
		}
	}
}

func TestSetMin(t *testing.T) {
	cases := []struct {
		want  Vector
		given Set
	}{
		{Vector{-1, -1}, NewSet(Vector{-1, 1}, Vector{1, -1})},
	}
	for _, c := range cases {
		got := c.given.Min()
		if got != c.want {
			t.Errorf("got %v want %v", got, c.want)
		}
	}
}

func TestSetMax(t *testing.T) {
	cases := []struct {
		want  Vector
		given Set
	}{
		{Vector{1, 1}, NewSet(Vector{-1, 1}, Vector{1, -1})},
	}
	for _, c := range cases {
		got := c.given.Max()
		if got != c.want {
			t.Errorf("got %v want %v", got, c.want)
		}
	}
}

func TestSetAdd(t *testing.T) {
	cases := []struct {
		want   Set
		set    Set
		vector Vector
	}{
		{
			NewSet(Vector{1, 2}, Vector{2, 4}),
			NewSet(Vector{0, 0}, Vector{1, 2}),
			Vector{1, 2},
		},
	}
	for _, c := range cases {
		got := c.set.Add(c.vector)
		if !got.Equal(c.want) {
			t.Errorf("got %#v want %#v", got, c.want)
		}
	}
}

func TestSetSub(t *testing.T) {
	cases := []struct {
		want   Set
		set    Set
		vector Vector
	}{
		{
			NewSet(Vector{0, 0}, Vector{1, 2}),
			NewSet(Vector{1, 2}, Vector{2, 4}),
			Vector{1, 2},
		},
	}
	for _, c := range cases {
		got := c.set.Sub(c.vector)
		if !got.Equal(c.want) {
			t.Errorf("got %v want %v", got, c.want)
		}
	}
}

func TestSetMul(t *testing.T) {
	cases := []struct {
		want   Set
		set    Set
		vector Vector
	}{
		{
			NewSet(Vector{-2, 1}, Vector{-4, 3}),
			NewSet(Vector{1, 2}, Vector{3, 4}),
			Vector{0, 1},
		},
	}
	for _, c := range cases {
		got := c.set.Mul(c.vector)
		if !got.Equal(c.want) {
			t.Errorf("got %v want %v", got, c.want)
		}
	}
}

func TestSetAppend(t *testing.T) {
	cases := []struct {
		want   Set
		set    Set
		vector Vector
	}{
		{
			NewSet(Vector{0, 0}, Vector{1, 1}),
			NewSet(Vector{0, 0}),
			Vector{1, 1},
		},
	}
	for _, c := range cases {
		got := c.set.Append(c.vector)
		if !got.Equal(c.want) {
			t.Errorf("got %v want %v", got, c.want)
		}
	}
}
