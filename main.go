package main

import (
	"log"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
	"fmt"
)

const (
	screenWidth  = 640
	screenHeight = 480
	ballsSpeed   = 4
	paddleSpeed  = 5 // The paddle will move 6 pixels per tick
)

type Object struct {
	X, Y, Width, Height int
}

type Paddle struct {
	Object
}

type Ball struct {
	Object
	dxdt int // x velocity per tick
	dxdy int // y velocity per tick
}

type Game struct {
	Paddle
	Ball
	score, highscore int
}

func main() {
	ebiten.SetWindowTitle("Pong Game in Ebitengine")
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := &Game{} // empty game struct
	err := ebiten.RunGame(g)
	if err != nil {
		log.Fatal(err)
	}

}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Rectangles representing the paddle and the ball
	vector.DrawFilledRect(screen, 
		float32(g.Paddle.X), float32(g.Paddle.Y), float32(g.Paddle.Width), float32(g.Paddle.Height), 
		color.White, false, 
	)
	vector.DrawFilledRect(screen, 
		float32(g.Ball.X), float32(g.Ball.Y), float32(g.Paddle.Width), float32(g.Paddle.Height), color.White,
		false,
	)

	scoreString := "Score: " + fmt.Sprint(g.score)
	text.Draw(screen, scoreString, basicfont.Face7x13, 10 ,10, color.White)
	highScoreString := "High Score: " + fmt.Sprint(g.highscore)
	text.Draw(screen, highScoreString, basicfont.Face7x13, 10, 30, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	 // Controls the size of our window, called every frame. We'll keep it constant.
	return screenWidth, screenHeight
}
