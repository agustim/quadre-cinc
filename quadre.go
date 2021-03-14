package main

import (
	"fmt"
)

const Height = 9
const Width = 9

type Quadre struct {
	lastNum int
	lastH   int
	lastW   int
	cela    [Height][Width]int
}

func (q *Quadre) Draw() {
	for h := 0; h < Height; h++ {
		for w := 0; w < Width; w++ {
			fmt.Printf("%2d ", q.cela[h][w])
		}
		fmt.Println()
	}
	fmt.Println()
}

func (src *Quadre) Copy() *Quadre {
	q := &Quadre{}
	q.lastNum = src.lastNum
	q.lastH = src.lastH
	q.lastW = src.lastW
	for h := 0; h < Height; h++ {
		for w := 0; w < Width; w++ {
			q.cela[h][w] = src.cela[h][w]
		}
	}
	return q
}

func Init() *Quadre {
	q := &Quadre{}
	for h := 0; h < Height; h++ {
		for w := 0; w < Width; w++ {
			q.cela[h][w] = 0
		}
	}
	q.lastH = 0
	q.lastW = 0
	return q
}

func (q *Quadre) First() {
	q.lastH = 0
	q.lastW = 0
	q.lastNum = 1
	q.cela[q.lastH][q.lastW] = q.lastNum
}

func Firsts() []Quadre {
	var qs []Quadre

	for h := 0; h < Height; h++ {
		for w := 0; w < Width; w++ {
			q := Init()
			q.lastH = h
			q.lastW = w
			q.lastNum = 1
			q.cela[q.lastH][q.lastW] = q.lastNum
			qs = append(qs, *q)
		}
	}
	return qs
}

func (q *Quadre) Final() bool {
	return q.lastNum == Height*Width
}

func (q *Quadre) Ciclic() bool {
	return (q.lastH-3 >= 0 && q.cela[q.lastH-3][q.lastW] == 1) ||
		(q.lastH+3 < Height && q.cela[q.lastH+3][q.lastW] == 1) ||
		(q.lastW-3 >= 0 && q.cela[q.lastH][q.lastW-3] == 1) ||
		(q.lastW+3 < Width && q.cela[q.lastH][q.lastW+3] == 1) ||
		(q.lastH-2 >= 0 && q.lastW-2 >= 0 && q.cela[q.lastH-2][q.lastW-2] == 1) ||
		(q.lastH-2 >= 0 && q.lastW+2 < Width && q.cela[q.lastH-2][q.lastW+2] == 1) ||
		(q.lastH+2 < Height && q.lastW-2 >= 0 && q.cela[q.lastH+2][q.lastW-2] == 1) ||
		(q.lastH+2 < Height && q.lastW+2 < Width && q.cela[q.lastH+2][q.lastW+2] == 1)
}

func (q *Quadre) WhereIsFirst() (int, int) {
	var h, w int
	for h = 0; h < Height; h++ {
		for w = 0; w < Width; w++ {
			if q.cela[h][w] == 1 {
				return h, w
			}
		}
	}
	return h, w
}

func (q *Quadre) Moure(h int, w int) *Quadre {
	if q.cela[h][w] != 0 {
		fmt.Printf("Error espai (%d,%d) ocupat per %d\n", h, w, q.cela[h][w])
	}
	num := q.lastNum
	num++
	q.cela[h][w] = num
	q.lastH = h
	q.lastW = w
	q.lastNum = num
	return q
}

func (q *Quadre) NextStep() []Quadre {
	var qs []Quadre

	// Check up
	if q.lastH-3 >= 0 && q.cela[q.lastH-3][q.lastW] == 0 {
		newq := q.Copy()
		newq.Moure(q.lastH-3, q.lastW)
		qs = append(qs, *newq)
	}
	// down
	if q.lastH+3 < Height && q.cela[q.lastH+3][q.lastW] == 0 {
		newq := q.Copy()
		newq.Moure(q.lastH+3, q.lastW)
		qs = append(qs, *newq)
	}
	// left
	if q.lastW-3 >= 0 && q.cela[q.lastH][q.lastW-3] == 0 {
		newq := q.Copy()
		newq.Moure(q.lastH, q.lastW-3)
		qs = append(qs, *newq)
	}
	// right
	if q.lastW+3 < Width && q.cela[q.lastH][q.lastW+3] == 0 {
		newq := q.Copy()
		newq.Moure(q.lastH, q.lastW+3)
		qs = append(qs, *newq)
	}
	// up-left
	if q.lastH-2 >= 0 && q.lastW-2 >= 0 && q.cela[q.lastH-2][q.lastW-2] == 0 {
		newq := q.Copy()
		newq.Moure(q.lastH-2, q.lastW-2)
		qs = append(qs, *newq)
	}
	// up-right
	if q.lastH-2 >= 0 && q.lastW+2 < Width && q.cela[q.lastH-2][q.lastW+2] == 0 {
		newq := q.Copy()
		newq.Moure(q.lastH-2, q.lastW+2)
		qs = append(qs, *newq)
	}
	// down-left
	if q.lastH+2 < Height && q.lastW-2 >= 0 && q.cela[q.lastH+2][q.lastW-2] == 0 {
		newq := q.Copy()
		newq.Moure(q.lastH+2, q.lastW-2)
		qs = append(qs, *newq)
	}
	// down-right
	if q.lastH+2 < Height && q.lastW+2 < Width && q.cela[q.lastH+2][q.lastW+2] == 0 {
		newq := q.Copy()
		newq.Moure(q.lastH+2, q.lastW+2)
		qs = append(qs, *newq)
	}
	return qs
}
