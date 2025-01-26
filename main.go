package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	SCREEN_WIDTH  = 640
	SCREEN_HEIGHT = 480
)

type CellType string
const (
	Empty CellType = "Empty"
	Down CellType = "Down"
	Across CellType = "Across"
	Grid CellType = "Grid"
)
type Cell struct {
	x int
	y int
	cellType CellType
}
type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func main() {
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}