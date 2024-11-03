package interfaces

type Plugin interface {
	ID() string
	Initialize() error
	Start() error
	Stop() error
}