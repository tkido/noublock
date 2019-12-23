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

func TestMsb(t *testing.T) {
	cases := []struct {
		want  int
		given uint64
	}{
		{-1, 0},
		{0, 1},
		{3, 8},
		{3, 9},
		{10, 1024},
		{63, 18446744073709551615},
	}
	for _, c := range cases {
		got := bits.Len64(c.given) - 1
		if got != c.want {
			t.Errorf("got %v want %v", got, c.want)
		}
	}
}
