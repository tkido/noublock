package main

import "fmt"

// Status is status of solver
type Status struct {
	Depth      int
	Board      uint64
	Candidates uint64
	Log        []int
}

// NewStatus returns new status
func NewStatus() Status {
	return Status{
		Depth:      0,
		Board:      ^uint64(0),
		Candidates: ^uint64(0),
		Log:        []int{},
	}
}

// Copy returns copied status
func (s Status) Copy() Status {
	newLog := []int{}
	for _, v := range s.Log {
		newLog = append(newLog, v)
	}
	return Status{
		Depth:      s.Depth,
		Board:      s.Board,
		Candidates: s.Candidates,
		Log:        newLog,
	}
}

// String returns string
func (s Status) String() string {
	return fmt.Sprintf("Depth: %d\n", s.Depth)
}
