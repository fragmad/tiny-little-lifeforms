package main

import (
	"math/rand"
	"time"
)

type cell struct {
	x     int
	y     int
	state bool
}

// redundent at the moment
func newCell(x, y int, state bool) cell {
	c := cell{x: x, y: y, state: true}
	return c
}

type board struct {
	width  int
	height int
	rows   [][]cell
}

func newBoard(w, h int, state bool) board {
	rows := make([][]cell, h, h)
	for y := 0; y < h; y++ {
		cols := make([]cell, w, w)
		for x := 0; x < w; x++ {
			c := cell{x, y, state}
			// fmt.Println(c)
			cols[x] = c
		}
		rows[y] = cols
		// fmt.Println("-------------")
	}

	b := board{w, h, rows}
	return b
}

func newRandBoard(w, h int) board {
	rows := make([][]cell, h, h)

	var src = rand.NewSource(time.Now().UnixNano())
	var r = rand.New(src)

	for y := 0; y < h; y++ {
		cols := make([]cell, w, w)
		for x := 0; x < w; x++ {
			state := r.Intn(2) != 0
			c := cell{x, y, state}
			// fmt.Println(c)
			cols[x] = c
		}
		rows[y] = cols
		// fmt.Println("-------------")
	}

	b := board{w, h, rows}
	return b
}

func decideCellNextState(c cell, b board) bool {
	cx := c.x
	cy := c.y
	cs := c.state
	new_state := false
	living_neighbours := 0

	type point struct {
		x int
		y int
	}

	possible_neighbours := make([]point, 8)
	n1 := point{cx, cy + 1}
	n2 := point{cx, cy - 1}
	n3 := point{cx + 1, cy}
	n4 := point{cx - 1, cy}
	n5 := point{cx + 1, cy + 1}
	n6 := point{cx - 1, cy - 1}
	n7 := point{cx - 1, cy + 1}
	n8 := point{cx - 1, cy + 1}

	possible_neighbours[0] = n1
	possible_neighbours[1] = n2
	possible_neighbours[2] = n3
	possible_neighbours[4] = n4
	possible_neighbours[4] = n5
	possible_neighbours[5] = n6
	possible_neighbours[6] = n7
	possible_neighbours[7] = n8

	// Either assume that cells of the grid are dead
	// or loop arond <-- try this first

	for _, n := range possible_neighbours {
		if n.x < 0 {
			// fmt.Println("1")
			n.x = n.x + b.width
		}
		if n.y < 0 {
			// fmt.Println("2")
			n.y = n.y + b.height
		}
		if n.x > b.width-1 {
			// fmt.Println("3")
			n.x = n.x - b.width
		}
		if n.y > b.height-1 {
			n.y = n.y - b.height
			// fmt.Println("4")
		}

		if b.rows[n.y][n.x].state == true {
			living_neighbours++
		} else {
			continue
		}

		if cs == true && living_neighbours < 2 {
			new_state = false
		} else if cs == true && (living_neighbours >= 2 && living_neighbours <= 3) {
			new_state = true
		} else if cs == true && living_neighbours > 3 {
			new_state = false
		} else if cs == false && living_neighbours > 3 {
			new_state = true
		}

	}
	return new_state
}
