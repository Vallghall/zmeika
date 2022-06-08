package game

import (
	"fmt"
	"github.com/Valghall/zmeika/internal/configs"
	"github.com/Valghall/zmeika/internal/food"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
)

type Mode int
type Direction int

const (
	ModeTitle Mode = iota
	ModeGame
	ModeGameOver
)

const (
	Left Direction = iota
	Up
	Right
	Down
)

var (
	headImage       *ebiten.Image
	foodImage       *ebiten.Image
	titleArcadeFont font.Face
	arcadeFont      font.Face
	smallArcadeFont font.Face
	titleTexts      []string
	texts           []string
)

func init() {
	foodImage = food.New()
	headImage = ebiten.NewImage(configs.ScreenWidth/20, configs.ScreenHeight/20)
	headImage.Fill(color.RGBA{0, 255, 0, 255})

	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	titleArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    configs.TitleFontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	arcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    configs.FontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	smallArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    configs.SmallFontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	headX         float64
	headY         float64
	foodPositionX float64
	foodPositionY float64

	direction Direction

	mode Mode

	bg    *ebiten.Image
	score int
}

func New() *Game {
	bg := ebiten.NewImage(configs.ScreenWidth, configs.ScreenHeight)
	bg.Fill(color.RGBA{R: 240, G: 150, B: 100, A: 1})

	return &Game{
		foodPositionX: float64(configs.ScreenWidth)/2 - 20,
		foodPositionY: float64(configs.ScreenHeight)/2 + 20,
		bg:            bg,
		direction:     Right,
		mode:          ModeTitle,
	}
}

func (g *Game) Update() error {
	if g.mode != ModeGame {
		if g.IsKeyPressed() {
			g.mode = ModeGame
		}
		return nil
	}
	g.ManageControlKey()
	dx, dy := g.Move()
	g.headX += dx
	g.headY += dy
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.bg, &ebiten.DrawImageOptions{})
	if g.mode != ModeTitle {
		g.drawHead(screen)
		g.drawFood(screen)
		return
	}
	switch g.mode {
	case ModeTitle:
		titleTexts = []string{"ZMEIKA"}
		texts = []string{"", "", "", "by", "Vallghall", "", "", "", "", "", "PRESS SPACE KEY", "", "OR A/B BUTTON", "", "OR TOUCH SCREEN"}
	case ModeGameOver:
		texts = []string{"", "GAME OVER!"}
	}
	for i, l := range titleTexts {
		x := (configs.ScreenWidth - len(l)*configs.TitleFontSize) / 2
		text.Draw(screen, l, titleArcadeFont, x, (i+4)*configs.TitleFontSize, color.RGBA{0, 255, 0, 255})
	}
	for i, l := range texts {
		x := (configs.ScreenWidth - len(l)*configs.FontSize) / 2
		text.Draw(screen, l, arcadeFont, x, (i+4)*configs.FontSize, color.RGBA{0, 255, 0, 255})
	}

	scoreStr := fmt.Sprintf("%04d", g.Score())
	text.Draw(screen, scoreStr, arcadeFont, configs.ScreenWidth-len(scoreStr)*configs.FontSize, configs.FontSize, color.RGBA{0, 255, 0, 255})
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return configs.ScreenWidth, configs.ScreenHeight
}

func (g Game) Score() int {
	return g.score
}

func (g *Game) IncrementScore() {
	g.score++
}

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
		if g.direction != Left {
			g.direction = Left
		}

		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) ||
		inpututil.IsKeyJustPressed(ebiten.KeyW) {
		if g.direction != Up {
			g.direction = Up
		}

		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) ||
		inpututil.IsKeyJustPressed(ebiten.KeyD) {
		if g.direction != Right {
			g.direction = Right
		}

		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) ||
		inpututil.IsKeyJustPressed(ebiten.KeyS) {
		if g.direction != Down {
			g.direction = Down
		}

		return
	}
}

func (g Game) Move() (dx, dy float64) {
	w, h := headImage.Size()

	switch g.direction {
	case Left:
		dx -= float64(w / 60)
	case Up:
		dy += float64(h / 60)
	case Right:
		dx += float64(w / 60)
	case Down:
		dy -= float64(w / 60)
	}
	return dx, dy
}

func (g *Game) drawHead(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{Filter: ebiten.FilterLinear}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("(%0.2f, %0.2f)", g.headX, g.headY))
	op.GeoM.Translate(g.headX, g.headY)
	op.GeoM.Translate(configs.HeadInitialX, configs.HeadInitialY)

	screen.DrawImage(headImage, op)
}

func (g Game) drawFood(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{Filter: ebiten.FilterLinear}

	op.GeoM.Translate(g.foodPositionX, g.foodPositionY)

	screen.DrawImage(foodImage, op)
}
