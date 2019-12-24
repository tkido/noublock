package main

import (
	"math/bits"
	"testing"
)

func TestCount(t *testing.T) {
	cases := []struct {
		want  int
		given uint64
	}{
		{0, 0},
		{1, 8},
		{2, 9},
		{64, 18446744073709551615},
	}
	for _, c := range cases {
		got := bits.OnesCount64(c.given)
		if got != c.want {
			t.Errorf("got %v want %v", got, c.want)
		}
	}
}

func TestFind(t *testing.T) {
	cases := []struct {
		want  int
		given uint64
	}{
		{64, 0},
		{0, 1},
		{1, 2},
		{10, 1024},
		{0, 18446744073709551615},
	}
	for _, c := range cases {
		got := find(c.given)
		if got != c.want {
			t.Errorf("got %v want %v", got, c.want)
		}
	}
}
