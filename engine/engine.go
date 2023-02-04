package engine

import (
	"math/rand"
)

type Coordinate struct {
	X int
	Y int
}

type Snake struct {
	Direction int
	Cells     []Coordinate
	Erase     Coordinate
	Alive     bool
}

type Game struct {
	SnakeObj Snake
	FoodPos  Coordinate
}

func GenNewCoord(X, Y int) (int, int) {
	f_x := rand.Intn(X)
	f_y := rand.Intn(Y)
	return f_x, f_y
}

func StartEngine(X, Y int) Game {
	x, y := GenNewCoord(X, Y)
	c := Coordinate{
		X: x,
		Y: y,
	}
	e := Coordinate{
		X: 0,
		Y: 0,
	}
	var cell_slice []Coordinate
	cell_slice = append(cell_slice, c)
	snake := Snake{
		Direction: 0,
		Cells:     cell_slice,
		Alive:     true,
		Erase:     e,
	}
	f_x, f_y := GenNewCoord(X, Y)
	f_c := Coordinate{
		X: f_x,
		Y: f_y,
	}
	game := Game{
		SnakeObj: snake,
		FoodPos:  f_c,
	}
	return game
}

func GenNewCell(game *Game, coord Coordinate) {
	game.SnakeObj.Cells = append(game.SnakeObj.Cells, coord)
}

func MoveSnake(game *Game, X, Y int) {
	e := game.SnakeObj.Cells[len(game.SnakeObj.Cells)-1]
	prev := game.SnakeObj.Cells[0]
	if game.SnakeObj.Direction == 108 {
		game.SnakeObj.Cells[0].X++
	} else if game.SnakeObj.Direction == 105 {
		game.SnakeObj.Cells[0].Y--
	} else if game.SnakeObj.Direction == 106 {
		game.SnakeObj.Cells[0].X--
	} else if game.SnakeObj.Direction == 107 {
		game.SnakeObj.Cells[0].Y++
	}

	for i := 1; i < len(game.SnakeObj.Cells); i++ {
		tmp := game.SnakeObj.Cells[i]
		game.SnakeObj.Cells[i] = prev
		prev = tmp
		if game.SnakeObj.Cells[0] == game.SnakeObj.Cells[i] {
			game.SnakeObj.Alive = false
		}
	}

	currX := game.SnakeObj.Cells[0].X
	currY := game.SnakeObj.Cells[0].Y
	if currX < 0 || currX >= X {
		game.SnakeObj.Alive = false
	}

	if currY < 0 || currY >= Y {
		game.SnakeObj.Alive = false
	}

	if currX == game.FoodPos.X && currY == game.FoodPos.Y {
		c := game.FoodPos
		f_x, f_y := GenNewCoord(X, Y)
		GenNewCell(game, c)
		game.FoodPos.X = f_x
		game.FoodPos.Y = f_y
	}
	game.SnakeObj.Erase = e
}
