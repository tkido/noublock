package main

import "fmt"

// Vector is simple vector
type Vector struct {
	X, Y int
}

func (v Vector) String() string {
	return fmt.Sprintf("(%d, %d)", v.X, v.Y)
}

// Add adds 2 vectors
func (v Vector) Add(o Vector) Vector {
	return Vector{v.X + o.X, v.Y + o.Y}
}

// Sub sub 2 vectors
func (v Vector) Sub(o Vector) Vector {
	return Vector{v.X - o.X, v.Y - o.Y}
}

// Mul mul 2 vectors
func (v Vector) Mul(o Vector) Vector {
	return Vector{v.X*o.X - v.Y*o.Y, v.X*o.Y - v.Y*o.X}
}
