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

// DeleteBody removes a body from the physics engine
func DeleteBody(b Body) {
	world.DestroyBody(b.b2body)
}

// Body represents a physical object in the simulation
type Body struct {
	b2body *box2d.B2Body
}

func (b *Body) Position() (x, y float64) {
	p := b.b2body.GetPosition()
	return p.X, p.Y
}

func (b *Body) PositionV() Vec2 {
	return Vec2(b.b2body.GetPosition())
}

func (b *Body) X() float64 {
	return b.b2body.GetPosition().X
}

func (b *Body) Y() float64 {
	return b.b2body.GetPosition().Y
}

func (b *Body) SetPosition(x, y float64) {
	angle := b.b2body.GetAngle()
	b.b2body.SetTransform(box2d.B2Vec2{X: x, Y: y}, angle)
}

func (b *Body) SetPositionV(v Vec2) {
	b.SetPosition(v.X, v.Y)
}

func (b *Body) SetAngle(a float64) {
	pos := b.b2body.GetPosition()
	b.b2body.SetTransform(pos, a)
}

func (b *Body) Impulse(impulse Vec2) {
	b.b2body.ApplyLinearImpulseToCenter(box2d.B2Vec2(impulse), true)
}

func (b *Body) ReachVelocity(vel Vec2) {
	m := b.b2body.GetMass()
	currVel := Vec2(b.b2body.GetLinearVelocity())
	vel.SubVec(currVel)
	vel.Scale(m)
	b.Impulse(vel)
}

func (b *Body) Force(force Vec2) {
	b.b2body.ApplyForceToCenter(force.toB2Vec2(), true)
}

// SetType allows to change the body type at runtime
func (b *Body) SetType(t BodyType) {
	switch t {
	case DynamicBody:
		b.b2body.SetType(box2d.B2BodyType.B2_dynamicBody)

	case KinematicBody:
		b.b2body.SetType(box2d.B2BodyType.B2_kinematicBody)

	case StaticBody:
		b.b2body.SetType(box2d.B2BodyType.B2_staticBody)
	}
}

func NewBody(opts ...BodyOpt) Body {
	// Default body definition
	def := box2d.MakeB2BodyDef()
	params := bodyParams{bd: def, fd: nil}

	// Apply optional body bodyDef
	for _, op := range opts {
		op(&params)
	}

	// Create a new body from its definition
	return Body{b2body: makeBody(params)}
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
