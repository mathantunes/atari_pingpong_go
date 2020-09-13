package image

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

func LoadGraphic(file string, r *sdl.Renderer) *sdl.Texture {
	sfc, err := sdl.LoadBMP(file)
	if err != nil {
		fmt.Println("Load Image failed", err)
		os.Exit(1)
	}
	defer sfc.Free()
	graphic, err := r.CreateTextureFromSurface(sfc)
	if err != nil {
		fmt.Println("CreateTextureFromSurface failed", err)
		os.Exit(1)
	}
	return graphic
}
