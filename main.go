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
	ballSpeed   = 4
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
	dydt int // y velocity per tick
}

type Game struct {
	paddle Paddle
	ball Ball
	score, highScore int
}

func main() {
	ebiten.SetWindowTitle("Pong Game in Ebitengine")
	ebiten.SetWindowSize(screenWidth, screenHeight)
	paddle := Paddle{
		Object: Object{
			X: 600, Y: 200, Width: 15, Height: 100,
		},
	}
	ball := Ball{
		Object: Object{
			X:0, Y:0, Width:15, Height:15,
		},
		dxdt: ballSpeed,
		dydt: ballSpeed,
	}

	g := &Game{
		paddle: paddle,
		ball: ball,
	} // empty game struct
	err := ebiten.RunGame(g)
	if err != nil {
		log.Fatal(err)
	}

}

func (g *Game) Update() error {
	g.paddle.MoveOnKeyPress()
	g.ball.Move()
	g.CollideWithWall()
	g.CollideWithPaddle()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Rectangles representing the paddle and the ball
	vector.DrawFilledRect(screen, 
		float32(g.paddle.X), float32(g.paddle.Y), float32(g.paddle.Width), float32(g.paddle.Height), 
		color.White, false, 
	)
	vector.DrawFilledRect(screen, 
		float32(g.ball.X), float32(g.ball.Y), float32(g.ball.Width), float32(g.ball.Height), color.White,
		false,
	)

	scoreString := "Score: " + fmt.Sprint(g.score)
	text.Draw(screen, scoreString, basicfont.Face7x13, 10 ,10, color.White)
	highScoreString := "High Score: " + fmt.Sprint(g.highScore)
	text.Draw(screen, highScoreString, basicfont.Face7x13, 10, 30, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	 // Controls the size of our window, called every frame. We'll keep it constant.
	return screenWidth, screenHeight
}

func (padd* Paddle) MoveOnKeyPress(){
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown){
		padd.Y += paddleSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp){
		padd.Y -= paddleSpeed
	}
}

func (ball* Ball) Move() { // The ball moves on its own
	/*
	for myself:

	Receiver (ball *Ball):
    This is the receiver of the method. It specifies that Move is a method for the type *Ball.
    *Ball means Move operates on a pointer to a Ball struct. This allows the method to modify
	 the actual instance of Ball that it is called on.
	*/
	ball.X += ball.dxdt
	ball.Y += ball.dydt
}

func(g *Game) Reset() {
	g.ball.X = 0
	g.ball.Y = 0

	g.score = 0
}

func (g *Game) CollideWithWall(){
	if g.ball.X >= screenWidth{
		g.Reset()
	}else if g.ball.X <=0{
		g.ball.dxdt = ballSpeed
	}else if g.ball.Y <=0{
		g.ball.dydt = ballSpeed
	}else if g.ball.Y >= screenHeight{
		g.ball.dydt = - ballSpeed
	}
}

func (g *Game) CollideWithPaddle(){
	if g.ball.X >= g.paddle.X && g.ball.Y>=g.paddle.Y && g.ball.Y <= g.paddle.Y + g.paddle.Height {
		g.ball.dxdt = -g.ball.dxdt
		g.score ++
		if g.score > g.highScore {
			g.highScore = g.score
		}
	}
}