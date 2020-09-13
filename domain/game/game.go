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
	ScoreBoard  *domain.ScoreBoard
	Status      Status
}

type Status int

const (
	Start Status = iota
	Playing
	GameOver
)

var aiCounter = 0

func New(pixels []byte, width, height int) *G {
	return &G{Pixels: pixels}
}
func (g *G) Init() {
	g.PaddleLeft = domain.NewPaddle(
		domain.NewPosition(50, 300),
		domain.NewSize(20, 100),
		domain.White,
		2000000,
	)
	g.PaddleRight = domain.NewPaddle(
		domain.NewPosition(750, 300),
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
	g.ScoreBoard = domain.NewScoreBoard(
		domain.NewPosition(200, 100),
		domain.NewPosition(600, 100),
		10,
		domain.White,
	)
	g.Ball.AddListener(g.ScoreBoard)
	g.Ball.AddListener(g)

	kbdDispatcher := keyboard.NewEventDispatcher()
	kbdDispatcher.AddListener(g.PaddleLeft)
	kbdDispatcher.AddListener(g)
	g.Pooler = keyboard.NewSDLEventPooler(kbdDispatcher)
	g.Status = Start
	g.UpdateFrame(0)
}

func (g *G) RunFrame(delta float32) {
	g.prepare(delta)
	if g.Status == Playing {
		g.UpdateFrame(delta)
	}
}

func (g *G) GetStatus() Status {
	return g.Status
}

func (g *G) prepare(delta float32) {
	Pools(g.Pooler)
	FrameRateCorrects(delta, g.Ball, g.PaddleLeft, g.PaddleRight)
}

func (g *G) UpdateFrame(delta float32) {
	if aiCounter > 10 {
		AIUpdates(g.Ball, g.PaddleRight)
		aiCounter = 0
	}
	aiCounter++
	Clear(g.Pixels)
	Updates(delta, g.Ball)
	Bounces(g.PaddleLeft, g.PaddleRight, g.Ball)
	Draws(g.Pixels, g.PaddleLeft, g.Ball, g.PaddleRight, g.ScoreBoard)
}
func (g *G) OnScore(evt domain.ScoreEvent) {
	g.Status = Start
	l, r, m := g.ScoreBoard.GetScore()
	if l > m {
		g.ScoreBoard.Reset()
		g.UpdateFrame(0)
		g.Status = GameOver
	} else if r > m {
		g.ScoreBoard.Reset()
		g.UpdateFrame(0)
		g.Status = GameOver
	}
}
func (g *G) Update(evt domain.KeyboardEvent) {
	if evt.Key == domain.Space && evt.Keydown > 0 {
		g.Status = Playing
	}
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
