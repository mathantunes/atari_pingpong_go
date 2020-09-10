package main

// Run go run main.go to test the installation of sdl

import (
	"fmt"

	"github.com/mathantunes/atari_pingpong_go/domain/game"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	// Width of screen
	Width = 800
	// Height of screen
	Height = 600
	// Delay SDL delay on game loop
	Delay = 16
)

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Println("Error Init: ", err)
		return
	}
	defer sdl.Quit()
	window, err := sdl.CreateWindow("Ping Pong - Golang", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, Width, Height, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println("Error create Window: ", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Error Create Renderer: ", err)
		return
	}

	defer renderer.Destroy()

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, Width, Height)
	if err != nil {
		fmt.Println("Error Create Texture: ", err)
		return
	}
	defer tex.Destroy()

	pixels := make([]byte, Width*Height*4)
	g := game.New(pixels, Width, Height)
	g.Init()
	for {
		g.GameLoop()
		tex.Update(nil, pixels, Width*4)
		renderer.Copy(tex, nil, nil)
		renderer.Present()
		sdl.Delay(Delay)
	}
}
