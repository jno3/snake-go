package input

import (
	"log"

	"game/engine"

	"github.com/mattn/go-tty"
)

func InputStart() *tty.TTY {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	return tty
}

func InputGet(game *engine.Game, tty *tty.TTY) {
	// tty, err := tty.Open()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer tty.Close()

	r, err := tty.ReadRune()
	if err != nil {
		log.Fatal(err)
	}

	if (r == 108) && (game.SnakeObj.Direction != 106) {
		game.SnakeObj.Direction = 108
	} else if (r == 105) && (game.SnakeObj.Direction != 107) {
		game.SnakeObj.Direction = 105
	} else if (r == 106) && (game.SnakeObj.Direction != 108) {
		game.SnakeObj.Direction = 106
	} else if (r == 107) && (game.SnakeObj.Direction != 105) {
		game.SnakeObj.Direction = 107
	}
}
