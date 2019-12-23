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
	count := 0
	pss := []PieceSet{}
	masks := []uint64{}
	for i, piece := range pieces {
		fmt.Printf("piece %d: %s\n", i, piece.String())
		mask := uint64(0)
		ps := PieceSet{piece}
		for j := 0; j < 3; j++ {
			piece = piece.Rotate()
			ps = ps.Append(piece)
		}
		piece = piece.Rotate().Mirror()
		ps = ps.Append(piece)
		for j := 0; j < 3; j++ {
			piece = piece.Rotate()
			ps = ps.Append(piece)
		}
		fmt.Printf("count = %d\n", len(ps))
		for j := range ps {
			mask |= 1 << (count + j)
		}
		for j, p := range ps {
			fmt.Printf("candidate %d\n", count+j)
			fmt.Println(p.Image())
			masks = append(masks, mask)
			printUint64("mask", mask)
		}
		pss = append(pss, ps)
		count += len(ps)
	}
	fmt.Println(len(pss))
	fmt.Println(count)

	board := uint64(0)
	candidates := ^uint64(0)

	printUint64("board", board)
	printUint64("candidates", candidates)
}

func printUint64(name string, i uint64) {
	fmt.Println(name)
	s := fmt.Sprintf("%064b", i)
	for j := 0; j < 8; j++ {
		fmt.Println(s[j*8 : (j+1)*8])
	}
	fmt.Println()
}
