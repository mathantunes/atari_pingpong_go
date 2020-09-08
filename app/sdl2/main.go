package main

// Run go run main.go to test the installation of sdl

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	// Width of screen
	Width = 800
	// Height of screen
	Height = 600
)

func main() {
	window, err := sdl.CreateWindow("SDL Test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, Width, Height, sdl.WINDOW_SHOWN)
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

	for y := 0; y <= Height; y++ {
		for x := 0; x < Width; x++ {
			setPixels(x, y, pixels)
		}
	}
	tex.Update(nil, pixels, Width*4)
	renderer.Copy(tex, nil, nil)
	renderer.Present()
	sdl.Delay(2000)
}

func setPixels(x, y int, pixels []byte) {
	pos := (y*Width + x) * 4
	r := byte(x % 255)
	g := byte(y % 255)
	if pos < len(pixels)-4 {
		pixels[pos] = r
		pixels[pos+1] = g
		pixels[pos+2] = 0
	}
}
