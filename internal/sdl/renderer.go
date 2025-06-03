package sdl

/*
#cgo pkg-config: sdl3
#include <SDL3/SDL.h>
*/
import "C"

type Color struct {
	R, G, B, A byte
}

type Renderer struct {
	renderer *C.SDL_Renderer
}

func (w *Window) CreateRenderer() *Renderer {
	renderer := C.SDL_CreateRenderer(w.handle, nil)
	if renderer == nil {
		panic(C.GoString(C.SDL_GetError()))
	}
	C.SDL_SetRenderScale(renderer, 2.0, 2.0)
	return &Renderer{renderer: renderer}
}

func (r *Renderer) Clear() {
	C.SDL_SetRenderDrawColor(r.renderer, 0, 0, 0, 255)
	C.SDL_RenderClear(r.renderer)
}

func (r *Renderer) DrawPixel(x, y int, color Color) {
	C.SDL_SetRenderDrawColor(r.renderer, C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.A))
	C.SDL_RenderPoint(r.renderer, C.float(x), C.float(y))
}

func (r *Renderer) Present() {
	C.SDL_RenderPresent(r.renderer)
}

func (r *Renderer) Destroy() {
	if r.renderer != nil {
		C.SDL_DestroyRenderer(r.renderer)
		r.renderer = nil
	}
}
