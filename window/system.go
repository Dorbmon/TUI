package window

import "runtime"

var fChan = make(chan func())

func init() {
	go func() {
		for {
			f := <-fChan
			runtime.LockOSThread()
			f()
			runtime.UnlockOSThread()
		}
	}()
}
func run(f func()) {
	fChan <- f
}
