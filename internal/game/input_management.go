package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g Game) IsKeyPressed() bool {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) ||
		inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		return true
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) ||
		inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return true
	}

	return false
}

func (g *Game) ManageControlKey() {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) ||
		inpututil.IsKeyJustPressed(ebiten.KeyA) {
		if g.direction != Right {
			g.direction = Left
		}

		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) ||
		inpututil.IsKeyJustPressed(ebiten.KeyW) {
		if g.direction != Down {
			g.direction = Up
		}

		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) ||
		inpututil.IsKeyJustPressed(ebiten.KeyD) {
		if g.direction != Left {
			g.direction = Right
		}

		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) ||
		inpututil.IsKeyJustPressed(ebiten.KeyS) {
		if g.direction != Up {
			g.direction = Down
		}

		return
	}
}

func (g *Game) manageExtraKeys() {
	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		g.secret = true
		secretPhrase = "LAG MOD ON BITCHES"
	}

}
