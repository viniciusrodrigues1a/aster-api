package controller

type CommandEmitter interface {
	Emit(msg interface{})
}
