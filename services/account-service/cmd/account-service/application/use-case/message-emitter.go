package usecase

type MessageEmitter interface {
	Emit(message interface{})
}
