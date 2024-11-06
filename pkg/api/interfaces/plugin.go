package interfaces

type Plugin interface {
	ID() string
	Dependencies() []string
	Capabilities() []string
	Initialize() error
	Start() error
	Stop() error
}
