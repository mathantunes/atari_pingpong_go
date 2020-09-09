package keyboard

import (
	"github.com/mathantunes/atari_pingpong_go/domain"
	"github.com/veandco/go-sdl2/sdl"
)

type EventDispatcher struct {
	sub []domain.KeyBoardListener
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{}
}

func (e *EventDispatcher) AddListener(l domain.KeyBoardListener) {
	e.sub = append(e.sub, l)
}

func (e *EventDispatcher) Dispatch(evt *sdl.KeyboardEvent) {
	for idx := range e.sub {
		e.sub[idx].Update(domain.KeyboardEvent{})
	}
}
