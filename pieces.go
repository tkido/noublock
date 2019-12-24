package main

// PieceSet is list of piece
type PieceSet []Set

// Append appends set to pieceSet
func (ps PieceSet) Append(s Set) PieceSet {
	for _, p := range ps {
		if p.Equal(s) {
			return ps
		}
	}
	ps = append(ps, s)
	return ps
}

var pieces = []Set{
	// ■■
	// ■■
	NewSet(
		Vector{0, 0},
		Vector{1, 0},
		Vector{0, 1},
		Vector{1, 1},
	),
	// 　■
	// ■■■
	// 　■
	NewSet(
		Vector{0, 0},
		Vector{-1, 1},
		Vector{0, 1},
		Vector{1, 1},
		Vector{0, 2},
	),
	// ■■■■■
	NewSet(
		Vector{0, 0},
		Vector{1, 0},
		Vector{2, 0},
		Vector{3, 0},
		Vector{4, 0},
	),
	// ■■■
	// ■　■
	NewSet(
		Vector{0, 0},
		Vector{1, 0},
		Vector{2, 0},
		Vector{0, 1},
		Vector{2, 1},
	),
	// ■■■
	// ■
	// ■
	NewSet(
		Vector{0, 0},
		Vector{1, 0},
		Vector{2, 0},
		Vector{0, 1},
		Vector{0, 2},
	),
	// 　■■
	// ■■
	// ■
	NewSet(
		Vector{0, 0},
		Vector{1, 0},
		Vector{-1, 1},
		Vector{0, 1},
		Vector{-1, 2},
	),
	// ■
	// ■■■
	// ■
	NewSet(
		Vector{0, 0},
		Vector{0, 1},
		Vector{1, 1},
		Vector{2, 1},
		Vector{0, 2},
	),
	// 　　■
	// ■■■
	// ■
	NewSet(
		Vector{0, 0},
		Vector{-2, 1},
		Vector{-1, 1},
		Vector{0, 1},
		Vector{-2, 2},
	),

	// ■■■■
	// ■
	NewSet(
		Vector{0, 0},
		Vector{1, 0},
		Vector{2, 0},
		Vector{3, 0},
		Vector{0, 1},
	),
	// ■■■■
	// 　■
	NewSet(
		Vector{0, 0},
		Vector{1, 0},
		Vector{2, 0},
		Vector{3, 0},
		Vector{1, 1},
	),
	// 　■■■
	// ■■
	NewSet(
		Vector{0, 0},
		Vector{1, 0},
		Vector{2, 0},
		Vector{-1, 1},
		Vector{0, 1},
	),
	// ■■■
	// ■■
	NewSet(
		Vector{0, 0},
		Vector{1, 0},
		Vector{2, 0},
		Vector{0, 1},
		Vector{1, 1},
	),
	// piece 10(a)
	// 　■
	// ■■■
	// ■
	NewSet(
		Vector{0, 0},
		Vector{-1, 1},
		Vector{0, 1},
		Vector{1, 1},
		Vector{-1, 2},
	),
}
