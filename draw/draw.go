package draw

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)
var IsGameEnd = false

var _clear map[string]func()

const Circle string = "o"
const Cross string = "x"

var IsPlayerTurn = true
var Board = make([][]string, 3)

func init() {
	FillZeroValues()

	_clear = make(map[string]func()) //Initialize it

	_clear[runtime.GOOS] = func() {
		var cmd *exec.Cmd

		switch pcOs := runtime.GOOS; pcOs {
		case "darwin", "linux":
			cmd = exec.Command("clear")
		case "windows":
			{
				cmd = exec.Command("cmd", "/c", "cls")
			}
		default:
			cmd = exec.Command("echo", pcOs+" platform is unsupported! I can't clear terminal screen :(")
		}

		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func FillZeroValues()  {
	// creating matrix slice
	for i := 0; i < len(Board); i++ {
		Board[i] = make([]string, 3)
		for k := 0; k < 3; k++ {
			Board[i][k] = "_"//strconv.Itoa(i + k + 1)
		}
	}
}

func getWinner() (bool, string) {
	shape := Circle

	for i := 0; i < 2; i++ {
		// я не помню когда писал хуже, но в 3 часа ночи
		// это показалось вполне правильным решением
		if shape == Board[0][0] &&
			shape == Board[0][1] &&
			shape == Board[0][2] || // first horizontal

			shape == Board[1][0] &&
				shape == Board[1][1] &&
				shape == Board[1][2] || // second horizontal

			shape == Board[2][0] &&
				shape == Board[2][1] &&
				shape == Board[2][2] || // third horizontal

			shape == Board[0][0] &&
				shape == Board[1][0] &&
				shape == Board[2][0] || // first vertical

			shape == Board[0][1] &&
				shape == Board[1][1] &&
				shape == Board[2][1] || // second vertical

			shape == Board[0][2] &&
				shape == Board[1][2] &&
				shape == Board[2][2] || // third vertical

			shape == Board[0][0] &&
				shape == Board[1][1] &&
				shape == Board[2][2] || // first diagonally

			shape == Board[0][2] &&
				shape == Board[1][1] &&
				shape == Board[2][0] { // second diagonally
			return true, shape
		}

		shape = Cross
	}

	return false, ""
}

func BoardPos(i int) (int,int) {
	var index int
	var subIndex int

	if i > 3 && i <= 6 { // second row
		index = 1
	} else if i > 6 && i <= 9 { // third row
		index = 2
	} // else will be default value

	subIndex = (i % 3) - 1

	if subIndex == -1 {
		subIndex = 2
	}

	return index, subIndex
}

func showNoMovesLeft()  {
	HighlightText("No any moves left. Game is end.")
	HighlightText("Press ESC to quit")
	HighlightText("or ENTER to restart")
}

func getIsNoMoves() bool {
	// game end
	gE := true

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if Board[i][j] != Circle && Board[i][j] != Cross {
				gE = false
				break
			}
		}
	}

	return gE
}

func Clear() {
	value, ok := _clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.

	if ok { //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("weird problem")
	}
}

func displayWinner(winner string) {
	HighlightText(winner + " won the game")
	HighlightText("Press ENTER to start again")
	HighlightText("Or ESC to stop the game")
}

func AppendShape(i int) {
	if IsGameEnd {
		showNoMovesLeft()
		return
	}

	index, subIndex := BoardPos(i)

	if Board[index][subIndex] == Circle || Board[index][subIndex] == Cross {
		HighlightText("You can't walk here")
		return
	}

	if IsPlayerTurn {
		Board[index][subIndex] = Circle
	} else {
		Board[index][subIndex] = Cross
	}

	IsPlayerTurn = !IsPlayerTurn

	RedrawMap()

	if winner, char := getWinner(); winner {
		IsGameEnd = true
		displayWinner(char)
		return
	} else {
		// is end
		if isGe := getIsNoMoves(); isGe {
			if isGe {
				IsGameEnd = true
				showNoMovesLeft()
				return
			}
		}
	}

	if !IsPlayerTurn {
		HighlightText("you (" + Circle + ") made a move. Now it's bot (" + Cross + ") turn")
	} else {
		HighlightText( "Bot made a move. Now it's your's (" + Circle + ") turn")
	}
}

func HighlightText(text string) {
	notice := color.New(color.Bold, color.FgGreen).PrintlnFunc()
	notice(text)
}

func RedrawMap() {
	Clear()

	// filling matrix slice with "press key" number
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			var curI int
			fmt.Print(" ")
			if i == 0 {
				curI = j + 1
			} else {
				curI = (i * 3) + (j + 1)
			}

			if _, err := strconv.Atoi(Board[i][j]); err == nil { //
				fmt.Print(curI)
			} else {
				fmt.Print(Board[i][j])
			}

			if j != 2 {
				fmt.Print(" |")
			}
		}

		if i == 0 || i == 1 {
			fmt.Println("\n-----------")
		} else {
			fmt.Println("\n ")
		}
	}
}
