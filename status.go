package main

// Status is status of solver
type Status struct {
	Board      uint64
	Candidates uint64
	Log        []int
}

// NewStatus returns new status
func NewStatus() Status {
	return Status{
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
		Board:      s.Board,
		Candidates: s.Candidates,
		Log:        newLog,
	}
}
