package infra

import (
	"github.com/mathantunes/atari_pingpong_go/domain"
	"github.com/veandco/go-sdl2/sdl"
)

type KeyboardDispatcher interface {
	Dispatch(*sdl.KeyboardEvent)
	AddListener(domain.KeyBoardListener)
}

type EventPooler interface {
	Pool()
}
