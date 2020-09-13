package game

import (
	"github.com/mathantunes/atari_pingpong_go/domain"
	"github.com/mathantunes/atari_pingpong_go/infra"
	"github.com/mathantunes/atari_pingpong_go/infra/keyboard"
)

type G struct {
	Pixels      []byte
	PaddleLeft  *domain.Paddle
	PaddleRight *domain.Paddle
	Ball        *domain.Ball
	Pooler      infra.EventPooler
}

var aiCounter = 0

func New(pixels []byte, width, height int) *G {
	return &G{Pixels: pixels}
}
func (g *G) Init() {
	g.PaddleLeft = domain.NewPaddle(
		domain.NewPosition(50, 100),
		domain.NewSize(20, 100),
		domain.White,
		2000000,
	)
	g.PaddleRight = domain.NewPaddle(
		domain.NewPosition(750, 100),
		domain.NewSize(20, 100),
		domain.White,
		2000000,
	)
	g.Ball = domain.NewBall(
		domain.NewPosition(300, 300),
		20,
		domain.White,
		domain.NewVelocity(300000, 300000),
	)

	kbdDispatcher := keyboard.NewEventDispatcher()
	kbdDispatcher.AddListener(g.PaddleLeft)
	g.Pooler = keyboard.NewSDLEventPooler(kbdDispatcher)
}

func (g *G) UpdateGame(delta float32) {
	Clear(g.Pixels)
	FrameRateCorrects(delta, g.Ball, g.PaddleLeft, g.PaddleRight)
	Pools(g.Pooler)

	if aiCounter > 2 {
		AIUpdates(g.Ball, g.PaddleRight)
		aiCounter = 0
	}
	aiCounter++
	Updates(delta, g.Ball)
	Bounces(g.PaddleLeft, g.PaddleRight, g.Ball)
	Draws(g.Pixels, g.PaddleLeft, g.Ball, g.PaddleRight)
}

func FrameRateCorrects(delta float32, frcs ...domain.FrameRateCorrect) {
	for _, frc := range frcs {
		frc.SetDelta(delta)
	}
}
func Pools(ps ...infra.EventPooler) {
	for _, p := range ps {
		p.Pool()
	}
}

func Draws(pxls []byte, drws ...domain.Drawable) {
	for _, d := range drws {
		d.Draw(pxls)
	}
}

func Updates(delta float32, ups ...domain.Updatable) {
	for _, u := range ups {
		u.Update()
	}
}

func AIUpdates(b *domain.Ball, ups ...domain.AutomaticPlayer) {
	for _, u := range ups {
		u.AutoUpdate(b)
	}
}

func Bounces(pLeft, pRight *domain.Paddle, bcs ...domain.Bouncer) {
	for _, b := range bcs {
		b.Bounce(pLeft, pRight)
	}
}

func Clear(pixels []byte) {
	for i := range pixels {
		pixels[i] = 0
	}
}
