package main

/* TODO:

- Implement a way to represent a cell -- DONE
- Implement a way to represent the board - DONE
- Implement a way to display the board in ASCII - DONE, but too far to long when combined with the above
- Implment a way to find out what the next state is - DONE
- Implement a way to update the board - DONE
- Divide project into seperate files - DONE
- Implement command line flags - DONE
- Imlement a way to input start states
- Implement gif output

- Assume X * X shape
*/

import (
	"flag"
	"fmt"
)

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
	return false
}

func main() {
	showDiagnostics := false
	runLife := true

	generationsPtr := flag.Int("generations", 10, "Number of generations to run")
	boardHeightPtr := flag.Int("height", 10, "Board height")
	boardWidthPtr := flag.Int("width", 10, "Board width")
	silenceOutputPtr := flag.Bool("silent", false, "Silence ASCII output")

	flag.Parse()

	numberOfGenerations := *generationsPtr
	silent := *silenceOutputPtr
	board_height := *boardHeightPtr
	board_width := *boardWidthPtr

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
		universe := newRandBoard(board_width, board_height)

		if !silent {
			fmt.Println("--- IT BEGINS --- ")
			fmt.Printf("A story of %d generations.\n", numberOfGenerations)
			textRenderBoard(&universe)
			fmt.Println("---")
		}

		for i := 1; i <= numberOfGenerations; i++ {
			if !silent {
				fmt.Println("--- Chapter", i, "--- ")
			}
			new_universe := nextGenerationBoard(universe)
			universe = new_universe
			if !silent {
				textRenderBoard(&new_universe)
			}
		}
		if !silent {
			fmt.Println("--- THE END --- ")
		}
	}
}
