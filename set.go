package main

import (
	"bytes"
	"reflect"
)

// Null is null struct
type Null struct{}

// Set is set of Vector
type Set map[Vector]Null

// NewSet returns new Set
func NewSet(vectors ...Vector) Set {
	s := Set{}
	for _, v := range vectors {
		s[v] = Null{}
	}
	return s
}

func (s Set) String() string {
	buf := bytes.Buffer{}
	buf.WriteString("Set{")
	for v := range s {
		buf.WriteString(v.String())
	}
	buf.WriteString("}")
	return buf.String()
}

// Equal returns bool
func (s Set) Equal(o Set) bool {
	return reflect.DeepEqual(s, o)
}

// Origin returns Vector
func (s Set) Origin() Vector {
	x, y := 64, 64
	for v := range s {
		if y > v.Y {
			x, y = v.X, v.Y
		} else if y == v.Y && x >= v.X {
			x, y = v.X, v.Y
		}
	}
	return Vector{x, y}
}

// Min returns Vector
func (s Set) Min() Vector {
	x, y := 64, 64
	for v := range s {
		if x > v.X {
			x = v.X
		}
		if y > v.Y {
			y = v.Y
		}
	}
	return Vector{x, y}
}

// Max returns Vector
func (s Set) Max() Vector {
	x, y := -64, -64
	for v := range s {
		if x < v.X {
			x = v.X
		}
		if y < v.Y {
			y = v.Y
		}
	}
	return Vector{x, y}
}

// Add move set
func (s Set) Add(v Vector) Set {
	vs := []Vector{}
	for k := range s {
		vs = append(vs, k.Add(v))
	}
	return NewSet(vs...)
}

// Sub move set
func (s Set) Sub(v Vector) Set {
	vs := []Vector{}
	for k := range s {
		vs = append(vs, k.Sub(v))
	}
	return NewSet(vs...)
}

// Mul move set
func (s Set) Mul(v Vector) Set {
	vs := []Vector{}
	for k := range s {
		vs = append(vs, k.Mul(v))
	}
	return NewSet(vs...)
}

// Rotate rotates set 90 degrees clockwise
func (s Set) Rotate() Set {
	s0 := s.Mul(Vector{0, 1})
	return s0.Sub(s0.Origin())
}

// Mirror returns mirrored set
func (s Set) Mirror() Set {
	vs := []Vector{}
	for k := range s {
		vs = append(vs, Vector{-k.X, k.Y})
	}
	m := NewSet(vs...)
	return m.Sub(m.Origin())
}

// Contain returns bool
func (s Set) Contain(v Vector) bool {
	_, ok := s[v]
	return ok
}

// Append appends vector to set
func (s Set) Append(v Vector) Set {
	s[v] = Null{}
	return s
}

// Image visualize set
func (s Set) Image() string {
	buf := bytes.Buffer{}
	for j := 0; j < 5; j++ {
		for i := -3; i < 5; i++ {
			char := '□'
			if s.Contain(Vector{i, j}) {
				char = '■'
			}
			buf.WriteRune(char)
		}
		buf.WriteRune('\n')
	}
	return buf.String()
}

// Board visualize board
func (s Set) Board() string {
	buf := bytes.Buffer{}
	for j := -1; j <= 8; j++ {
		for i := -1; i <= 8; i++ {
			char := '□'
			if s.Contain(Vector{i, j}) {
				char = '■'
			}
			buf.WriteRune(char)
		}
		buf.WriteRune('\n')
	}
	return buf.String()
}
