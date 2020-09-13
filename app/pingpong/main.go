package main

// Run go run main.go to test the installation of sdl

import (
	"fmt"
	"time"

	"github.com/mathantunes/atari_pingpong_go/domain/game"
	"github.com/mathantunes/atari_pingpong_go/infra/image"
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

	gameOverTex := image.LoadGraphic("./assets/gameover.bmp", renderer)
	defer gameOverTex.Destroy()
	_, _, w, h, err := gameOverTex.Query()
	if err != nil {
		fmt.Println("gameOverTex Query error: ", err)
		return
	}
	pixels := make([]byte, Width*Height*4)
	g := game.New(pixels, Width, Height)
	g.Init()
	var Maxfps uint32 = 200
	var delay uint32 = 1000 / Maxfps
	var frameDur float32 = 0
	for {
		frameStart := time.Now()
		g.RunFrame(frameDur)
		if st := g.GetStatus(); st == game.GameOver {
			renderer.Copy(gameOverTex, &sdl.Rect{0, 0, w, h}, &sdl.Rect{Width/2 - w/2, Height/2 - h/2, w / 2, h / 2})
		} else {
			tex.Update(nil, pixels, Width*4)
			renderer.Copy(tex, nil, nil)
		}
		renderer.Present()
		frameDur = float32(time.Since(frameStart).Seconds()) / 1000
		if frameDur < 0.005 {
			sdl.Delay(delay)
			frameDur = float32(time.Since(frameStart).Seconds()) / 1000
		}
	}
}
