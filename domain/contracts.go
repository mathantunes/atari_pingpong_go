package domain

type GameLoop interface {
	GameLoop()
}

type Drawable interface {
	Draw([]byte)
}

type Updatable interface {
	Update()
}

type KeyBoardListener interface {
	Update(KeyboardEvent)
}

type AutomaticPlayer interface {
	AutoUpdate(*Ball)
}

type Bouncer interface {
	Bounce(*Paddle, *Paddle)
}
