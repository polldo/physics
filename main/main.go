package main

import (
	"image/color"

	"body"

	"github.com/go-p5/p5"
)

const (
	width  = 640
	height = 640

	rad = 5.
)

var (
	ballBody body.Body
)

func setup() {
	p5.Canvas(width, height)
	p5.Stroke(nil)
	p5.Fill(color.RGBA{R: 255, A: 204})

	body.SetGravity(0.3, 8)
	ballBody = body.NewBody(body.Type(body.DynamicBody), body.Circle(rad))
	ballBody.SetPosition(100, 100)
}

func main() {
	p5.Run(setup, draw)
}

func draw() {
	body.Update()
	p5.Background(color.Gray{Y: 220})

	p5.Fill(color.RGBA{R: 255, A: 204})
	p5.Ellipse(ballBody.X(), ballBody.Y(), 2*rad, 2*rad)
}
