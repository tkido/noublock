package main

import (
	"fmt"
	"math/bits"
)

var (
	availables = []uint64{}
	areaSets   = [][]uint64{}
	masks      = []uint64{}
)

func main() {
	count := 0
	shapes := PieceSet{}
	for i, piece := range pieces {
		fmt.Printf("piece %d: %s\n", i, piece.String())
		mask := uint64(0)
		ps := PieceSet{piece}
		shapes = shapes.Append(piece)
		for j := 0; j < 3; j++ {
			piece = piece.Rotate()
			ps = ps.Append(piece)
			shapes = shapes.Append(piece)
		}
		piece = piece.Rotate().Mirror()
		ps = ps.Append(piece)
		shapes = shapes.Append(piece)
		for j := 0; j < 3; j++ {
			piece = piece.Rotate()
			ps = ps.Append(piece)
			shapes = shapes.Append(piece)
		}
		// fmt.Printf("count = %d\n", len(ps))
		for j := range ps {
			mask |= 1 << (count + j)
		}
		for j, p := range ps {
			fmt.Printf("candidate %d\n", count+j)
			fmt.Println(p.Image())
			masks = append(masks, mask)
			// printUint64("mask", mask)
		}
		count += len(ps)
	}
	fmt.Println(count)

	// for i := 0; i < 64; i++ {
	// 	fmt.Printf("shape %d\n", i)
	// 	fmt.Println(shapes[i].Image())
	// 	name := fmt.Sprintf("mask: %d", i)
	// 	printUint64(name, masks[i])
	// }

	for j := 0; j < 8; j++ {
		for i := 0; i < 8; i++ {
			v := Vector{i, j}
			available := uint64(0)
			areas := []uint64{}
		LABEL1:
			for n := 0; n < 64; n++ {
				s := shapes[n]
				s = s.Add(v)
				area := uint64(0)
				for k := range s {
					if k.X < 0 || 8 <= k.X || k.Y < 0 || 8 <= k.Y {
						areas = append(areas, uint64(0))
						continue LABEL1
					}
					area |= 1 << (k.Y*8 + k.X)
				}
				// printUint64("area", area)
				areas = append(areas, area)
				available |= 1 << n
			}
			// fmt.Printf("%sに置けるのは\n", Vector{i, j}.String())
			// printUint64("available", available)
			availables = append(availables, available)
			areaSets = append(areaSets, areas)
		}
	}

	// printUint64("availables[0]", availables[0])
	// printUint64("areaSets[0][0]", areaSets[0][0])
	// printUint64("areaSets[0][2]", areaSets[0][2])
	// printUint64("areaSets[1][1]", areaSets[1][1])

	loop(NewStatus())
	fmt.Printf("clearCount = %d\n", clearCount/8)
}

func find(i uint64) int {
	return bits.TrailingZeros64(i)
}

var clearCount = 0

func loop(s Status) bool {
	// printUint64("s.Board", s.Board)
	// printUint64("s.Candidates", s.Candidates)
	// fmt.Println(s.Log)

	target := find(s.Board)
	// fmt.Printf("target = %d\n", target)

	// fmt.Println(target)
	if target == 64 {
		// fmt.Println("Cleared!!")
		fmt.Println(s.Log)
		clearCount++
		return false
	}
	candidates := s.Candidates & availables[target]
	// printUint64("s.Candidates", s.Candidates)
	// printUint64("availables[target]", availables[target])
	// printUint64("candidates", candidates)

	for {
		candidate := find(candidates)
		// fmt.Printf("candidate = %d\n", candidate)
		if candidate == 64 {
			// fmt.Println("Impossible!!")
			return false
		}
		area := areaSets[target][candidate]
		if area&s.Board == area {
			copied := s.Copy()
			copied.Board = copied.Board &^ area
			copied.Candidates = copied.Candidates &^ masks[candidate]
			copied.Log = append(copied.Log, candidate)
			if loop(copied) {
				return true
			}
		}
		// fmt.Printf("collision(%d)(%d)\n", target, candidate)
		candidates = candidates &^ (1 << candidate)
	}
}

func printUint64(name string, i uint64) {
	fmt.Println(name)
	s := fmt.Sprintf("%064b", bits.Reverse64(i))
	for j := 0; j < 8; j++ {
		fmt.Println(s[j*8 : (j+1)*8])
	}
	fmt.Println()
}
