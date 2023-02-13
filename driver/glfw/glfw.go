package driverGLFW

import (
	"runtime"
	"sync"
	"time"

	drivers "github.com/dorbmon/TUI/driver"
	"github.com/dorbmon/TUI/window"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var glfwInit sync.Once

type GLFWDriver struct {
	runOnMain chan func()
	done      chan int
	windows   []window.Window
}

func New() (drivers.Driver, error) {
	var err error
	glfwInit.Do(func() {
		err = glfw.Init()
		if err != nil {
			return
		}
	})
	if err != nil {
		return nil, err
	}
	return &GLFWDriver{
		runOnMain: make(chan func()),
		done:      make(chan int),
		windows:   make([]window.Window, 0),
	}, nil
}
func (g *GLFWDriver) Run() error {
	runtime.LockOSThread()
	eventTick := time.NewTicker(time.Second / 60)
	for {
		select {
		case <-g.done:
			// clear and exit
			eventTick.Stop()
			return nil
		case <-eventTick.C:
			glfw.PollEvents()
			newWin := make([]window.Window, 0)
			for _, rwin := range g.windows {
				win := rwin.(*GWindow)
				if win.window == nil {
					continue
				}
				if win.window.ShouldClose() {
					win.window.Destroy()
					continue
				}
				win.viewLock.RLock()
				shouldResize := win.shouldResize
				newSizeWidth := win.resizeWidth
				newSizeHeight := win.resizeHeight
				fullScreen := win.fullScreen
				win.viewLock.RUnlock()
				if shouldResize && !fullScreen {
					// then do the resize
					win.viewLock.Lock()
					shouldResize = win.shouldResize
					win.shouldResize = false
					win.viewLock.Unlock()
					if shouldResize {
						win.window.SetSize(newSizeWidth, newSizeHeight)
					}
				}
				newWin = append(newWin, rwin)
			}
			g.windows = newWin
			if len(g.windows) == 0 {
				// there is no window exists... then exit the app
				g.Quit()
			}
		}

	}
}
func (g *GLFWDriver) drawThread() {

}
func (g *GLFWDriver) Quit() {

}
