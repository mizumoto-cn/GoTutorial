// It's easy to add new levels
// But difficult to add new types
package abstractfactory

type Programmer interface {
	WriteCode() string
}

type Architect interface {
	Design() string
}

type AbstractFactory interface {
	CreateProgrammer() Programmer
	CreateArchitect() Architect
}

type FrontendProgrammer struct{}

func (p *FrontendProgrammer) WriteCode() string {
	return "FrontendProgrammer writes html and css with DreamWeaver"
}

type BackendProgrammer struct{}

func (p *BackendProgrammer) WriteCode() string {
	return "BackendProgrammer writes golang with vscode and vim-plugin"
}

type FrontendArchitect struct{}

func (a *FrontendArchitect) Design() string {
	return "FrontendArchitect designs html and css with Adobe XD"
}

type BackendArchitect struct{}

func (a *BackendArchitect) Design() string {
	return "BackendArchitect designs microservices and try to deploy them with kubernetes"
}

type FrontendFactory struct{}

func (f *FrontendFactory) CreateProgrammer() Programmer {
	return &FrontendProgrammer{}
}

func (f *FrontendFactory) CreateArchitect() Architect {
	return &FrontendArchitect{}
}

type BackendFactory struct{}

func (f *BackendFactory) CreateProgrammer() Programmer {
	return &BackendProgrammer{}
}

func (f *BackendFactory) CreateArchitect() Architect {
	return &BackendArchitect{}
}
