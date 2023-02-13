package tui

type Driver interface {
	Run() error // this is called every tick
}
