package main

var pieces = []Set{
	// piece 0
	// ■■
	// ■■
	NewSet(
		Vector{0, 0},
		Vector{1, 0},
		Vector{0, 1},
		Vector{1, 1},
	),
	// piece 1
	// ■■■■■
	NewSet(
		Vector{0, 0},
		Vector{1, 0},
		Vector{2, 0},
		Vector{3, 0},
		Vector{4, 0},
	),
	// piece 2
	// ■■■■
	// ■
	NewSet(
		Vector{0, 0},
		Vector{1, 0},
		Vector{2, 0},
		Vector{3, 0},
		Vector{0, 1},
	),
	// piece 3
	// ■■■■
	// 　■
	NewSet(
		Vector{0, 0},
		Vector{1, 0},
		Vector{2, 0},
		Vector{3, 0},
		Vector{1, 1},
	),
	// piece 4
	// 　■■■
	// ■■
	NewSet(
		Vector{0, 0},
		Vector{1, 0},
		Vector{2, 0},
		Vector{-1, 1},
		Vector{0, 1},
	),
	// piece 5
	// ■■■
	// ■■
	NewSet(
		Vector{0, 0},
		Vector{1, 0},
		Vector{2, 0},
		Vector{0, 1},
		Vector{1, 1},
	),
	// piece 6
	// ■■■
	// ■　■
	NewSet(
		Vector{0, 0},
		Vector{1, 0},
		Vector{2, 0},
		Vector{0, 1},
		Vector{2, 1},
	),
	// piece 7
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
	// piece 8
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
	// piece 9
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
	// piece 11(b)
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
	// piece 12(c)
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
}
