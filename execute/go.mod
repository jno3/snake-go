module game/execute

go 1.20

replace game/draw => ../draw

replace game/engine => ../engine

replace game/input => ../input

require (
	game/draw v0.0.0-00010101000000-000000000000
	game/engine v0.0.0-00010101000000-000000000000
	game/input v0.0.0-00010101000000-000000000000
)

require (
	github.com/mattn/go-isatty v0.0.10 // indirect
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/mattn/go-tty v0.0.4 // indirect
	github.com/nsf/termbox-go v1.1.1 // indirect
	golang.org/x/sys v0.0.0-20191120155948-bd437916bb0e // indirect
)
