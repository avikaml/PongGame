package main

import(
	"log"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth = 640
	screenHeight = 480
	ballsSpeed = 4
	paddleSpeed = 5 // The paddle will move 6 pixels per tick
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



