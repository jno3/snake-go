package draw

import (
	"game/engine"
	"time"

	termbox "github.com/nsf/termbox-go"
)

func DrawBoard(X, Y int) {
	termbox.Init()
	for i := 0; i <= Y; i++ {
		termbox.SetCell(X, i, '#', termbox.ColorDefault, termbox.ColorRed)
	}
	for i := 0; i <= X; i++ {
		termbox.SetCell(i, Y, '#', termbox.ColorDefault, termbox.ColorRed)
	}
	termbox.Flush()
}

func DrawSnake(game *engine.Game, ms int) {
	// length := len(engine.Cells)

	for i := range game.SnakeObj.Cells {
		X := game.SnakeObj.Cells[i].X
		Y := game.SnakeObj.Cells[i].Y
		termbox.SetCell(X, Y, 'O', termbox.ColorGreen, termbox.ColorLightYellow)
		if game.SnakeObj.Direction != 0 {
			eX := game.SnakeObj.Erase.X
			eY := game.SnakeObj.Erase.Y
			termbox.SetCell(eX, eY, ' ', termbox.ColorDefault, termbox.ColorDefault)
		}
	}
	fX := game.FoodPos.X
	fY := game.FoodPos.Y
	termbox.SetCell(fX, fY, ' ', termbox.ColorDefault, termbox.ColorGreen)
	termbox.Flush()

	if !game.SnakeObj.Alive {
		termbox.Close()
	}
	time.Sleep(time.Millisecond * time.Duration(ms))
}
