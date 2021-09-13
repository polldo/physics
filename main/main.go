package main

import (
	"image/color"

	"body"

	"github.com/go-p5/p5"
	"github.com/polldo/box2d"
)

const (
	width  = 640
	height = 640

	ratio = 0.005

	rad = 5.
)

var (
	ballBody *box2d.B2Body
)

func setup() {
	p5.Canvas(width, height)
	p5.Stroke(nil)
	p5.Fill(color.RGBA{R: 255, A: 204})

	body.SetGravity(body.V2(0.3, 0.3))
	ballBody = body.NewBody(body.Type(body.DynamicBody), body.Circle(rad))
	ballBody.SetTransform(box2d.B2Vec2{100, 100}, 0)
}

func main() {
	p5.Run(setup, draw)
}

func draw() {
	body.Update()
	p5.Background(color.Gray{Y: 220})

	p5.Fill(color.RGBA{R: 255, A: 204})
	p5.Ellipse(ballBody.GetPosition().X, ballBody.GetPosition().Y, 2*rad, 2*rad)
}
