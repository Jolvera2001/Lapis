package kernel

type Plugin interface {
	ID() string
	Dependencies() []string
	Capabilities() []string
	Initialize(api API) error
	Start() error
	Stop() error
}
