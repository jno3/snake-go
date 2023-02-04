package main

import (
	"game/draw"
	"game/engine"
	"game/input"
)

const (
	CANVAS_Y = 30
	CANVAS_X = 10
	SPEED    = 150
)

func main() {
	game := engine.StartEngine(CANVAS_Y, CANVAS_X)
	draw.DrawBoard(CANVAS_Y, CANVAS_X)
	tty := input.InputStart()
	for {
		draw.DrawSnake(&game, SPEED)
		engine.MoveSnake(&game, CANVAS_Y, CANVAS_X)
		go input.InputGet(&game, tty)
		if !game.SnakeObj.Alive {
			draw.DrawSnake(&game, SPEED)
			break
		}
	}
}
