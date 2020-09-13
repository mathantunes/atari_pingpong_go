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

type ScoreDispatcher interface {
	Dispatch(ScoreEvent)
	AddListener(ScoreListener)
}

type ScoreListener interface {
	OnScore(ScoreEvent)
}

type FrameRateCorrect interface {
	SetDelta(float32)
}
