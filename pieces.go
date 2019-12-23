package main

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
