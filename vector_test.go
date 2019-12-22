package main

import "testing"

func TestVectorEqual(t *testing.T) {
	cases := []struct {
		want   bool
		v1, v2 Vector
	}{
		{true, Vector{1, 1}, Vector{1, 1}},
		{true, Vector{2, 1}, Vector{2, 1}},
		{true, Vector{1, 2}, Vector{1, 2}},
		{false, Vector{1, 1}, Vector{1, 2}},
		{false, Vector{1, 1}, Vector{2, 1}},
		{false, Vector{1, 1}, Vector{2, 2}},
		{true, Vector{-1, -1}, Vector{-1, -1}},
		{false, Vector{-1, -1}, Vector{-2, -1}},
		{false, Vector{-1, -1}, Vector{-1, -2}},
	}
	for _, c := range cases {
		got := c.v1 == c.v2
		if got != c.want {
			t.Errorf("got %v want %v", got, c.want)
		}
	}
}

func TestVectorAdd(t *testing.T) {
	cases := []struct {
		want, v1, v2 Vector
	}{
		{Vector{2, 2}, Vector{1, 1}, Vector{1, 1}},
		{Vector{0, 0}, Vector{-1, 1}, Vector{1, -1}},
		{Vector{-2, -2}, Vector{-1, -1}, Vector{-1, -1}},
	}
	for _, c := range cases {
		got := c.v1.Add(c.v2)
		if got != c.want {
			t.Errorf("got %v want %v", got, c.want)
		}
	}
}

func TestVectorSub(t *testing.T) {
	cases := []struct {
		want, v1, v2 Vector
	}{
		{Vector{0, 0}, Vector{1, 1}, Vector{1, 1}},
		{Vector{-2, 2}, Vector{-1, 1}, Vector{1, -1}},
		{Vector{0, 0}, Vector{-1, -1}, Vector{-1, -1}},
	}
	for _, c := range cases {
		got := c.v1.Sub(c.v2)
		if got != c.want {
			t.Errorf("got %v want %v", got, c.want)
		}
	}
}

func TestVectorMul(t *testing.T) {
	cases := []struct {
		want, v1, v2 Vector
	}{
		{Vector{0, 1}, Vector{1, 0}, Vector{0, 1}},
		{Vector{-1, 0}, Vector{0, 1}, Vector{0, 1}},
		{Vector{0, -1}, Vector{-1, 0}, Vector{0, 1}},
		{Vector{1, 0}, Vector{0, -1}, Vector{0, 1}},
	}
	for _, c := range cases {
		got := c.v1.Mul(c.v2)
		if got != c.want {
			t.Errorf("got %v want %v", got, c.want)
		}
	}
}
