package main

/* TODO:

- Implement a way to represent a cell -- DONE
- Implement a way to represent the board - DONE
- Implement a way to display the board in ASCII - DONE, but too far to long when combined with the above
- Implment a way to find out what the next state is - DONE
- Implement a way to update the board - DONE
- Implement command line flags
- Imlement a way to input start states
- Implement gif output

- Assume X * X shape
*/

import (
	"fmt"
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

func nextGenerationBoard(current_board board) board {
	var new_board board = current_board

	for i := 0; i < current_board.width; i++ {
		for j := 0; j < current_board.height; j++ {
			new_cell := cell{new_board.rows[i][j].x,
				new_board.rows[i][j].y,
				decideCellNextState(current_board.rows[i][j], current_board)}

			new_board.rows[i][j] = new_cell
		}
	}
	return new_board
}

func textRenderBoardDiagnostic(b *board) {
	for i := 0; i < b.width; i++ {
		for j := 0; j < b.height; j++ {
			fmt.Print(b.rows[i])
			// fmt.Print(b.rows[i][j].x, b.rows[i][j].y)
		}
		fmt.Println("")
	}
}

func textRenderBoard(b *board) {
	for i := 0; i < b.width; i++ {
		for j := 0; j < b.height; j++ {
			if b.rows[i][j].state == true {
				fmt.Print("X")
			} else if b.rows[i][j].state == false {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func compareBoards(b1 *board, b2 *board) bool {

}

func main() {
	showDiagnostics := false
	runLife := true
	numberOfGenerations := 10000

	if showDiagnostics == true {
		// Cells
		// c1 := newCell(1, 1, true)
		// c2 := cell{x: 1, y: 1, state: true}
		// fmt.Println(c1)
		// fmt.Println(c2)
		// fmt.Println(c1 == c2)

		// // Boards
		// b1 := newBoard(10, 10, false)
		b2 := newBoard(10, 10, true)

		// fmt.Println("---- Boards ----- ")
		// fmt.Println("---- Board 1 ----- ")
		// textRenderBoardDiagnostic(&b1)
		// textRenderBoard(&b1)
		// fmt.Println("---- Board 2 ----- ")
		textRenderBoard(&b2)
		// textRenderBoardDiagnostic(&b2)

		fmt.Println("--------")
		// fmt.Println(b2.rows[0])
		// fmt.Println(b2.rows[0][1])

		// fmt.Println("-----")
		// fmt.Println(decideCellNextState(b2.rows[0][0], b2))
		next_gen := nextGenerationBoard(b2)
		textRenderBoard(&next_gen)
	}

	if runLife == true {
		// universe := newBoard(10, 10, true)
		universe := newRandBoard(100, 100)

		fmt.Println("--- IT BEGINS --- ")
		fmt.Printf("A story of %d generations.", numberOfGenerations)
		textRenderBoard(&universe)
		fmt.Println("---")
		for i := 1; i <= numberOfGenerations; i++ {
			fmt.Println("--- Chapter", i, "--- ")
			new_universe := nextGenerationBoard(universe)
			universe = new_universe
			textRenderBoard(&new_universe)
			fmt.Println("--- ")
		}
		fmt.Println("--- THE END --- ")
	}
}
