module game/draw

go 1.20

require (
	game/engine v0.0.0-00010101000000-000000000000
	github.com/nsf/termbox-go v1.1.1
)

require github.com/mattn/go-runewidth v0.0.9 // indirect

replace game/engine => ../engine
