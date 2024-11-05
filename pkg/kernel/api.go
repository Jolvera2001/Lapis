package kernel

type API interface {
	Initialize() error
	Shutdown() error
	Version() string
}

type CoreAPI struct {
	version string
}

func NewCoreAPI() *CoreAPI {
	return &CoreAPI{
		version: "0.0.1",
	}
}

func (c *CoreAPI) Initialize() error {
	return nil
}

func (c *CoreAPI) Shutdown() error {
	return nil
}

func (c *CoreAPI) Version() string {
	return c.version
}
