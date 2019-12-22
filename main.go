package main

import "fmt"

type PieceSet []Set

func (ps PieceSet) Append(s Set) PieceSet {
	for _, p := range ps {
		if p.Equal(s) {
			return ps
		}
	}
	ps = append(ps, s)
	return ps
}

func main() {
	pss := []PieceSet{}
	for i, piece := range pieces {
		fmt.Printf("piece %d: %s\n", i, piece.String())
		ps := PieceSet{}
		for j := 0; j < 4; j++ {
			piece = piece.Rotate()
			ps = ps.Append(piece)
		}
		piece = piece.Mirror()
		for j := 0; j < 4; j++ {
			piece = piece.Rotate()
			ps = ps.Append(piece)
		}
		fmt.Printf("count = %d\n", len(ps))
		for _, p := range ps {
			fmt.Println(p.Image())
		}
		pss = append(pss, ps)
	}

	board := Set{}
	for i := -1; i <= 8; i++ {
		board[Vector{i, -1}] = Null{}
		board[Vector{i, 8}] = Null{}
	}
	for j := -1; j <= 8; j++ {
		board[Vector{-1, j}] = Null{}
		board[Vector{8, j}] = Null{}
	}
	fmt.Println(board.Board())

}
