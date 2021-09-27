package bot

import (
	"math"
	"tictactoe/draw"
	"time"
)

type bestMove struct {
	row int
	col int
}

// This function returns true if there are moves
// remaining on the board. It returns false if
// there are no moves left to play.
func isMovesLeft(board [][]string) bool {
	// Package bot algorithm taken & rewrote to GO from
	// https://www.geeksforgeeks.org/minimax-algorithm-in-game-theory-set-3-tic-tac-toe-ai-finding-optimal-move/
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == "_" {
				return true
			}
		}
	}

	return false
}

// This is evaluate ion function as discussed
// in the previous article ( http://goo.gl/sJgv68 )
func evaluate(b [][]string) int {

	// Checking for Rows for X or O victory.
	for row := 0; row < 3; row++ {
		if b[row][0] == b[row][1] && b[row][1] == b[row][2] {
			if b[row][0] == draw.Cross {
				return +10
			} else if b[row][0] == draw.Circle {
				return -10
			}
		}
	}

	// Checking for Columns for X or O victory.
	for col := 0; col < 3; col++ {
		if b[0][col] == b[1][col] && b[1][col] == b[2][col] {
			if b[0][col] == draw.Cross {
				return +10
			} else if b[0][col] == draw.Circle {
				return -10
			}
		}
	}

	// Checking for Diagonals for X or O victory.
	if b[0][0] == b[1][1] && b[1][1] == b[2][2] {
		if b[0][0] == draw.Cross {
			return +10
		} else if b[0][0] == draw.Circle {
			return -10
		}
	}

	if b[0][2] == b[1][1] && b[1][1] == b[2][0] {
		if b[0][2] == draw.Cross {
			return +10
		} else if b[0][2] == draw.Circle {
			return -10
		}
	}

	// Else if none of them have
	// won then return 0
	return 0
}

// This is the minimax function. It
// considers all the possible ways
// the game can go and returns the
// value of the board
func minimax(board [][]string, depth int, isMax bool) int {
	score := evaluate(board)

	// If Maximizer has won the game
	// return his/her evaluated score
	if score == 10 {
		return score
	}

	// If Minimizer has won the game
	// return his/her evaluated score
	if score == -10 {
		return score
	}

	// If there are no more moves and
	// no winner then it is a tie
	if isMovesLeft(board) == false {
		return 0
	}

	// If this maximizer's move
	if isMax {
		best := -1000

		// Traverse all cells
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				// Check if cell is empty
				if board[i][j] == "_" {
					// Make the move
					board[i][j] = draw.Cross

					// Call minimax recursively
					// and choose the maximum value
					best = int(math.Max(float64(best), float64(minimax(board, depth+1, !isMax))))

					// Undo the move
					board[i][j] = "_"
				}
			}
		}

		return best
	} else { // If this minimizer's move else
		best := 1000

		// Traverse all cells
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {

				// Check if cell is empty
				if board[i][j] == "_" {

					// Make the move
					board[i][j] = draw.Circle

					// Call minimax recursively and
					// choose the minimum value
					best = int(math.Min(float64(best), float64(minimax(board, depth+1, !isMax))))

					// Undo the move
					board[i][j] = "_"
				}
			}
		}

		return best
	}
}

// findBestMove This will return the best possible
// move for the draw.Cross
func findBestMove(board [][]string) bestMove {
	bestVal := -1000
	var bs bestMove
	bs.row = -1
	bs.col = -1

	// Traverse all cells, evaluate
	// minimax function for all empty
	// cells. And return the cell
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// Check if cell is empty
			if board[i][j] == "_" {
				// Make the move
				board[i][j] = draw.Cross

				// compute evaluation function
				// for this move.
				moveVal := minimax(board, 0, false)

				// Undo the move
				board[i][j] = "_"

				// If the value of the current move
				// is more than the best value, then
				// update best
				if moveVal > bestVal {
					bs.row = i
					bs.col = j
					bestVal = moveVal
				}
			}
		}
	}

	//document.write("The value of the best Move " +
	//"is : ", bestVal + "<br><br>");

	return bs
}

func Move() {
	bs := findBestMove(draw.Board)
	time.Sleep(500 * time.Millisecond)
	i := 0
	switch row := bs.row; row {
	case 0:
		{
			i = 0
		}
	case 1:
		{
			i = 3
		}
	case 2:
		{
			i = 6
		}
	}

	switch col := bs.col; col {
	case 0:
		{
			i += 1
		}
	case 1:
		{
			i += 2
		}
	case 2:
		{
			i += 3
		}
	}

	// bot doesn't know best move. value is -1. Game is Over
	if i < 1 {
		return
	}

	draw.AppendShape(i)
}
