package main

import (
	"image/color"

	"github.com/polldo/physics"

	"github.com/go-p5/p5"
)

const (
	width  = 640
	height = 640

	rad = 5.
)

var (
	ballBody physics.Body
)

func setup() {
	p5.Canvas(width, height)
	p5.Stroke(nil)
	p5.Fill(color.RGBA{R: 255, A: 204})

	physics.SetGravity(0.3, 8)
	ballBody = physics.NewBody(physics.Type(physics.DynamicBody), physics.Circle(rad))
	ballBody.SetPosition(100, 100)
}

func main() {
	p5.Run(setup, draw)
}

func draw() {
	physics.Update()
	p5.Background(color.Gray{Y: 220})

	p5.Fill(color.RGBA{R: 255, A: 204})
	p5.Ellipse(ballBody.X(), ballBody.Y(), 2*rad, 2*rad)
}
