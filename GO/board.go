package main

import "math/rand"

type Board struct {
	cells      [][]int
	xdim, ydim int
}

func (b *Board) Get(x, y int) int {
	if x < 0 || y < 0 || x >= b.xdim || y >= b.ydim {
		return 0
	}
	return b.cells[y][x]
}

func (b *Board) Set(x, y int, c int) {
	b.cells[y][x] = c
}

func (b *Board) String() string {
	msg := ""
	alive := '■'
	dead := '□'
	for _, x := range b.cells {
		for _, v := range x {
			val := dead
			if v == 1 {
				val = alive
			}
			msg += string(val)
		}
		msg += "\n"
	}
	return msg
}

func (b *Board) Neighbours(x, y int) int {
	xs := []int{x - 1, x, x + 1}
	ys := []int{y - 1, y, y + 1}
	n := 0
	for _, X := range xs {
		for _, Y := range ys {
			if b.Get(X, Y) == 1 && !(X == x && Y == y) {
				n++
			}
		}
	}
	return n
}

func (b *Board) IsAlive(x, y int) int {
	n := b.Neighbours(x, y)
	switch val := b.Get(x, y); {
	case val == 1 && !(n == 2 || n == 3):
		return 0
	case val == 0 && n == 3:
		return 1
	default:
		return val
	}

}

func (b *Board) Copy() Board {
	copy_cells := make([][]int, len(b.cells))
	for i, v := range b.cells {
		copy_cells[i] = make([]int, len(v))
		copy(copy_cells[i], v)
	}
	return Board{copy_cells, b.xdim, b.ydim}
}

func (b *Board) NextPhase() {
	bef_b := b.Copy()
	for x := 0; x < b.xdim; x++ {
		for y := 0; y < b.ydim; y++ {
			b.Set(x, y, bef_b.IsAlive(x, y))
		}
	}
}

func NewBoard(x, y int) Board {
	c := make([][]int, y)
	for i := 0; i < y; i++ {
		c[i] = make([]int, x)
	}

	return Board{c, x, y}
}

func RandBoard(x, y int, prob float64) Board {
	b := NewBoard(x, y)
	for j, v := range b.cells {
		for i := range v {
			val := 0
			if rand.Float64() < prob {
				val = 1
			}
			b.Set(i, j, val)
		}
	}
	return b
}

func (b *Board) changeCell(x, y int) {
	b.Set(x, y, 1-b.Get(x, y))
}
