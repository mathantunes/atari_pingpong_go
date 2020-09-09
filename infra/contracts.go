package infra

import "github.com/veandco/go-sdl2/sdl"

type KeyboardDispatcher interface {
	Dispatch(*sdl.KeyboardEvent)
}
