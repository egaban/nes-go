package sdl

/*
#cgo pkg-config: sdl3
#include <SDL3/SDL.h>
*/
import "C"
import "unsafe"

const (
	INIT_VIDEO = C.SDL_INIT_VIDEO
)

type Event int

const (
	QUIT Event = iota
)

type Window struct {
	handle *C.SDL_Window
}

func Init() {
	if !C.SDL_Init(INIT_VIDEO) {
		panic(C.GoString(C.SDL_GetError()))
	}
}

func CreateWindow() *Window {
	window := C.SDL_CreateWindow(
		C.CString("NES Emulator"),
		1024,
		480,
		0,
	)

	if window == nil {
		panic(C.GoString(C.SDL_GetError()))
	}

	return &Window{handle: window}
}

func (w *Window) Destroy() {
	C.SDL_DestroyWindow(w.handle)
	w.handle = nil
}

func (w *Window) PollEvents() []Event {
	var event C.SDL_Event
	result := make([]Event, 0)

	for C.SDL_PollEvent(&event) {
		eventType := *(*C.Uint32)(unsafe.Pointer(&event))

		if eventType == C.SDL_EVENT_QUIT {
			switch eventType {
			case C.SDL_EVENT_QUIT:
				result = append(result, QUIT)
			default:
				panic("Unknown event type")
			}
		}
	}

	return result
}
