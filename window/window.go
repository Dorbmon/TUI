package window

type Window interface {
	Resize(width, height int)
	Run() error
}
