package keyboard

import (
	"os"

	"github.com/mathantunes/atari_pingpong_go/domain"
	"github.com/mathantunes/atari_pingpong_go/infra"
	"github.com/veandco/go-sdl2/sdl"
)

type SDLEventPooler struct {
	Dispatcher infra.KeyboardDispatcher
}

func NewSDLEventPooler(dispatcher infra.KeyboardDispatcher) *SDLEventPooler {
	return &SDLEventPooler{dispatcher}
}

func (s *SDLEventPooler) Pool() {
	for evt := sdl.PollEvent(); evt != nil; evt = sdl.PollEvent() {
		switch t := evt.(type) {
		case *sdl.QuitEvent:
			os.Exit(0)
		case *sdl.KeyboardEvent:
			s.Dispatcher.Dispatch(t)
		}
	}
}

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
		e.sub[idx].Update(domain.KeyboardEvent{
			Key:     toKey(evt),
			Keydown: evt.State,
		})
	}
}

func toKey(evt *sdl.KeyboardEvent) domain.Key {
	switch evt.Keysym.Sym {
	case sdl.K_UP:
		return domain.ArrowUp
	case sdl.K_DOWN:
		return domain.ArrowDown
	default:
		return domain.Key(999)
	}
}
