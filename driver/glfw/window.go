package driverGLFW

import (
	"runtime"
	"sync"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type GWindow struct {
	window *glfw.Window

	viewLock     sync.RWMutex
	shouldResize bool
	resizeWidth  int
	resizeHeight int
	fullScreen   bool
	visible      bool
}

func NewGWindow(width, height int, title string) (*GWindow, error) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	ret := &GWindow{
		shouldResize: false,
		visible:      true,
	}
	var err error
	if err != nil {
		return nil, err
	}
	return ret, nil
}
func (w *GWindow) Run() error {
	for !w.window.ShouldClose() {
		w.window.SwapBuffers()
	}
	return nil
}
func (w *GWindow) Resize(width, height int) {
	w.viewLock.Lock()
	defer w.viewLock.Unlock()
	w.shouldResize = true
	w.resizeHeight = height
	w.resizeWidth = width
}
