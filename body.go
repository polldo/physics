package body

import (
	"github.com/polldo/box2d"
)

type BodyType uint8

const (
	DynamicBody BodyType = iota
	KinematicBody
	StaticBody
)

type bodyParams struct {
	bd box2d.B2BodyDef
	fd *box2d.B2FixtureDef
}
type BodyOpt func(*bodyParams)

func DeleteBody(b *box2d.B2Body) {
	world.DestroyBody(b)
}

func NewBody(opts ...BodyOpt) *box2d.B2Body {
	// Default body definition
	def := box2d.MakeB2BodyDef()
	params := bodyParams{bd: def, fd: nil}

	// Apply optional body bodyDef
	for _, op := range opts {
		op(&params)
	}

	// Create a new body from its definition
	return makeBody(params)
}

func makeBody(params bodyParams) *box2d.B2Body {
	body := world.CreateBody(&params.bd)
	if params.fd != nil {
		body.CreateFixtureFromDef(params.fd)
	}
	return body
}

func Type(t BodyType) BodyOpt {
	return func(p *bodyParams) {
		switch t {
		case DynamicBody:
			p.bd.Type = box2d.B2BodyType.B2_dynamicBody

		case KinematicBody:
			p.bd.Type = box2d.B2BodyType.B2_kinematicBody

		case StaticBody:
			p.bd.Type = box2d.B2BodyType.B2_staticBody

		}
	}
}

func Damping(linear float64, angular float64) BodyOpt {
	return func(p *bodyParams) {
		p.bd.LinearDamping = linear
		p.bd.AngularDamping = angular
	}
}

func AllowSleep() BodyOpt {
	return func(p *bodyParams) {
		p.bd.AllowSleep = true
	}
}

func FixRotation() BodyOpt {
	return func(p *bodyParams) {
		p.bd.FixedRotation = true
	}
}

func Bullet() BodyOpt {
	return func(p *bodyParams) {
		p.bd.Bullet = true
	}
}

func ScaleGravity(scale float64) BodyOpt {
	return func(p *bodyParams) {
		p.bd.GravityScale = scale
	}
}

func Circle(radius float64) BodyOpt {
	return func(p *bodyParams) {
		c := box2d.NewB2CircleShape()
		c.SetRadius(radius)
		fd := box2d.MakeB2FixtureDef()
		fd.Shape = c
	}
}

func Edge(v1, v2 box2d.B2Vec2) BodyOpt {
	return func(p *bodyParams) {
		c := box2d.NewB2EdgeShape()
		c.Set(v1, v2)
		fd := box2d.MakeB2FixtureDef()
		fd.Shape = c
	}
}

func Polygon(vertices []box2d.B2Vec2) BodyOpt {
	return func(p *bodyParams) {
		c := box2d.NewB2PolygonShape()
		c.Set(vertices, len(vertices))
		fd := box2d.MakeB2FixtureDef()
		fd.Shape = c
	}
}

func Rectangle(base, height float64) BodyOpt {
	return func(p *bodyParams) {
		c := box2d.NewB2PolygonShape()
		c.SetAsBox(base/2, height/2)
		fd := box2d.MakeB2FixtureDef()
		fd.Shape = c
	}
}
