// Observer pattern
//  in which an object, named the subject, maintains a list of its dependents
// called observers, and notifies them automatically of any changes,
// usually by calling a method on each dependent.
package observer

type Observer struct {
	data string
}

type IObserver interface {
	Notify(string)
}

type Subject struct {
	Array []IObserver
}

type ISubject interface {
	Add(observer ...IObserver)
	Notify(string)
}

func (o *Observer) Notify(s string) {
	o.data = s
}

func (s *Subject) Add(observer ...IObserver) {
	s.Array = append(s.Array, observer...)
}

func (s *Subject) Notify(str string) {
	for _, o := range s.Array {
		o.Notify(str)
	}
}
