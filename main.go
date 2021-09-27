package main

import (
	"github.com/eiannone/keyboard"
	"strconv"
	"tictactoe/bot"
	"tictactoe/draw"
)

func listenUserInput() {
	keysEvents, err := keyboard.GetKeys(10)

	if err != nil {
		panic(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		event := <-keysEvents
		if event.Err != nil {
			panic(event.Err)
		}

		if event.Key == keyboard.KeyEsc {
			break
		}

		if event.Key == keyboard.KeyEnter {
			startNewGame()
		}

		switch key := event.Rune; key {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			{
				if draw.IsPlayerTurn {
					if v, err := strconv.Atoi(string(event.Rune)); err == nil {
						//player turn
						draw.AppendShape(v)
						// bot turn
						bot.Move()
					}
				}
			}
		default:
			{
				draw.HighlightText("This button does nothing.")
			}
		}
	}
}

func startNewGame() {
	draw.FillZeroValues()
	draw.Clear()
	draw.IsGameEnd = false
	draw.RedrawMap()
	bot.Move()
	draw.HighlightText("Welcome to tic-tac-toe terminal game")
	draw.HighlightText("In front of your terminal with cells and rows")
	draw.HighlightText("To start the game, press the number key on your keyboard (not numpad)")
	draw.HighlightText(draw.Cross + " (bot) always goes first (i'm lazy to rewrite it, sorry)")
	draw.HighlightText("Press ESC to quit")
	draw.HighlightText("or ENTER to restart")
}

func main() {
	startNewGame()
	listenUserInput()
}
