package domain

// Ball Represents the ping pong ball
type Ball struct {
	pos    Position
	radius float32
	vel    Velocity
	color  Color
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

// Update ball position and velocity
func (b *Ball) Update() {
	b.pos.x += b.vel.x
	b.pos.y += b.vel.y
	b.tryBounceOffLimits()
}

// Bounce on collision with paddle
func (b *Ball) Bounce(pLeft *Paddle, pRigth *Paddle) {

	if int(b.pos.x-b.radius) < int(pLeft.pos.x+pLeft.size.w/2) {
		if int(b.pos.y) > int(pLeft.pos.y-pLeft.size.h/2) && int(b.pos.y) < int(pLeft.pos.y+pLeft.size.h/2) {
			b.vel.x *= -1
		}
	}
	if int(b.pos.x+b.radius) > int(pRigth.pos.x-pRigth.size.w/2) {
		if int(b.pos.y) > int(pRigth.pos.y-pRigth.size.h/2) && int(b.pos.y) < int(pRigth.pos.y+pRigth.size.h/2) {
			b.vel.x *= -1
		}
	}
}

func (b *Ball) tryBounceOffLimits() {
	fRadius := float32(b.radius)
	if b.pos.y-fRadius < 0 || b.pos.y+fRadius > 600 {
		b.vel.y *= -1
	}
	if b.pos.x-fRadius < 0 || b.pos.x+fRadius > 800 {
		b.pos.x = 300
		b.pos.y = 300
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
	pos   Position
	size  Size
	color Color
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
		p.pos.y -= 20
	} else if evt.Key == ArrowDown && evt.Keydown > 0 {
		p.pos.y += 20
	}
}

// AutoUpdate automatic paddle
func (p *Paddle) AutoUpdate(b *Ball) {
	// p.pos.y = b.pos.y
	if p.pos.y > b.pos.y {
		p.pos.y -= 20
	}
	if p.pos.y < b.pos.y {
		p.pos.y += 20
	}
}

// NewPaddle Initialize paddle
func NewPaddle(p Position, s Size, c Color) *Paddle {
	return &Paddle{
		pos:   p,
		size:  s,
		color: c,
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

const (
	// Width of screen
	Width = 800
	// Height of screen
	Height = 600
)

var (
	// White RGB Color
	White = Color{255, 255, 255}
)
