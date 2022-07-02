package simpleFactory

type Factory struct{}

type Product interface {
	create() string
}

func (f *Factory) Create(name string) Product {
	switch name {
	case "pizza":
		return &Pizza{}
	case "steak":
		return &Steak{}
	}
	return nil
}

type Pizza struct{}

func (p *Pizza) create() string {
	return "cooked pizza"
}

type Steak struct{}

func (s *Steak) create() string {
	return "cooked steak"
}
