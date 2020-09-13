package domain

// Ball Represents the ping pong ball
type Ball struct {
	pos    Position
	radius float32
	vel    Velocity
	color  Color
	delta  float32
	subs   []ScoreListener
}

// Draw ball
func (b *Ball) Draw(pixels []byte) {
	// Iterate over a square, in each pixel, we check if inside or outside the radius
	for y := -b.radius; y < b.radius; y++ { //Iterate on Y first for best memory management
		for x := -b.radius; x < b.radius; x++ {
			if x*x+y*y < b.radius*b.radius {
				setPixels(int(b.pos.x+x), int(b.pos.y+y), b.color, pixels)
			}
		}
	}
}

// SetDelta stores frame rate delta on ball instance
func (b *Ball) SetDelta(delta float32) {
	b.delta = delta
}

// Update ball position and velocity
func (b *Ball) Update() {
	b.pos.x += b.vel.x * b.delta
	b.pos.y += b.vel.y * b.delta
	b.tryBounceOffLimits()
}

// Bounce on collision with paddle
func (b *Ball) Bounce(pLeft *Paddle, pRigth *Paddle) {

	if int(b.pos.x-b.radius) < int(pLeft.pos.x+pLeft.size.w/2) {
		if int(b.pos.y) > int(pLeft.pos.y-pLeft.size.h/2) && int(b.pos.y) < int(pLeft.pos.y+pLeft.size.h/2) {
			b.vel.x *= -1
			b.pos.x = pLeft.pos.x + pLeft.size.w/2 + b.radius
		}
	}
	if int(b.pos.x+b.radius) > int(pRigth.pos.x-pRigth.size.w/2) {
		if int(b.pos.y) > int(pRigth.pos.y-pRigth.size.h/2) && int(b.pos.y) < int(pRigth.pos.y+pRigth.size.h/2) {
			b.vel.x *= -1
			b.pos.x = pRigth.pos.x - pRigth.size.w/2 - b.radius
		}
	}
}

func (b *Ball) tryBounceOffLimits() {
	fRadius := float32(b.radius)
	if b.pos.y-fRadius < 0 || b.pos.y+fRadius > 600 {
		b.vel.y *= -1
	}
	if b.pos.x-fRadius < 0 {
		b.pos.x = Width / 2
		b.pos.y = Height / 2
		b.Dispatch(ScoreEvent{Right: true})
	} else if b.pos.x+fRadius > 800 {
		b.pos.x = Width / 2
		b.pos.y = Height / 2
		b.Dispatch(ScoreEvent{Left: true})
	}
}

func (b *Ball) AddListener(l ScoreListener) {
	b.subs = append(b.subs, l)
}

func (b *Ball) Dispatch(evt ScoreEvent) {
	for _, l := range b.subs {
		l.OnScore(evt)
	}
}

// NewBall Initializes a ball
func NewBall(p Position, r float32, c Color, v Velocity) *Ball {
	return &Ball{
		pos:    p,
		radius: r,
		color:  c,
		vel:    v,
	}
}

// Paddle Represents the players paddle
type Paddle struct {
	pos    Position
	size   Size
	color  Color
	delta  float32
	ySpeed float32
}

// SetDelta stores frame rate delta on paddle instance
func (p *Paddle) SetDelta(delta float32) {
	p.delta = delta
}

// Draw implementation of Drawable for Paddle
func (p *Paddle) Draw(pixels []byte) {
	st := NewPosition(p.pos.x-float32(p.size.w)/2.0, p.pos.y-float32(p.size.h)/2.0)
	for y := 0; y < int(p.size.h); y++ { //Iterate on Y first for best memory management
		for x := 0; x < int(p.size.w); x++ {
			setPixels(int(st.x)+x, int(st.y)+y, p.color, pixels)
		}
	}
}

// Update paddle position
func (p *Paddle) Update(evt KeyboardEvent) {
	if evt.Key == ArrowUp && evt.Keydown > 0 {
		p.pos.y -= p.ySpeed * p.delta
	} else if evt.Key == ArrowDown && evt.Keydown > 0 {
		p.pos.y += p.ySpeed * p.delta
	}
}

// AutoUpdate automatic paddle
func (p *Paddle) AutoUpdate(b *Ball) {
	// p.pos.y = b.pos.y
	if p.pos.y > b.pos.y {
		p.pos.y -= p.ySpeed * p.delta
	}
	if p.pos.y < b.pos.y {
		p.pos.y += p.ySpeed * p.delta
	}
}

// NewPaddle Initialize paddle
func NewPaddle(p Position, s Size, c Color, ySpeed float32) *Paddle {
	return &Paddle{
		pos:    p,
		size:   s,
		color:  c,
		ySpeed: ySpeed,
	}
}

// Position represents a position on x and y axis
type Position struct {
	x, y float32
}

// NewPosition Initialize position
func NewPosition(x, y float32) Position {
	return Position{x, y}
}

// Velocity represents a velocity on x and y axis
type Velocity struct {
	x, y float32
}

// NewVelocity Initialize velocity
func NewVelocity(x, y float32) Velocity {
	return Velocity{x, y}
}

// Color Represents a byte color using RGB
type Color struct {
	r, g, b byte
}

// Size represents a Height and Width
type Size struct {
	w, h float32
}

// NewSize Initialize size
func NewSize(w, h float32) Size {
	return Size{w, h}
}

func setPixels(x, y int, c Color, pixels []byte) {
	pos := (y*Width + x) * 4
	if pos < len(pixels)-4 && pos >= 0 {
		pixels[pos] = c.r
		pixels[pos+1] = c.g
		pixels[pos+2] = c.b
	}
}

type ScoreBoard struct {
	lPos       Position
	rPos       Position
	numberSize float32
	color      Color
	Left       int
	Right      int
	MaxScore   int
}

func NewScoreBoard(lPos, rPos Position, numberSize float32, color Color) *ScoreBoard {
	return &ScoreBoard{
		lPos:       lPos,
		rPos:       rPos,
		numberSize: numberSize,
		color:      color,
		MaxScore:   3,
	}
}

func (sc *ScoreBoard) OnScore(evt ScoreEvent) {
	if evt.Left {
		sc.Left++
	} else if evt.Right {
		sc.Right++
	}
}

func (sc *ScoreBoard) GetScore() (left, right, max int) {
	left = sc.Left
	right = sc.Right
	max = sc.MaxScore
	return
}
func (sc *ScoreBoard) Reset() {
	sc.Left = 0
	sc.Right = 0
}

func (sc *ScoreBoard) Draw(pxls []byte) {
	if sc.Left > sc.MaxScore {

	}
	if sc.Right > sc.MaxScore {

	}
	sc.drawNumber(sc.lPos, sc.Left, pxls)
	sc.drawNumber(sc.rPos, sc.Right, pxls)
}

func (sc *ScoreBoard) drawNumber(pos Position, score int, pxls []byte) {
	startX := pos.x - sc.numberSize*3/2
	startY := pos.y - sc.numberSize*5/2
	numberArr, ok := PossibleScores[score]
	if !ok {

	}
	for idx, n := range numberArr {
		if n == 1 {
			for y := startY; y < startY+sc.numberSize; y++ {
				for x := startX; x < startX+sc.numberSize; x++ {
					setPixels(int(x), int(y), sc.color, pxls)
				}
			}
		}
		startX += sc.numberSize
		if (idx+1)%3 == 0 {
			startY += sc.numberSize
			startX -= sc.numberSize * 3
		}
	}
}
