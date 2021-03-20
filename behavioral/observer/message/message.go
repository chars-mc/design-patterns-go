package message

import "github.com/chars-mc/design-patterns-go/behavioral/observer/interfaces"

type Message struct {
	Message   string
	observers map[string]interfaces.Observer
}

func (m *Message) AddObserver(name string, o interfaces.Observer) {
	if m.observers == nil {
		m.observers = make(map[string]interfaces.Observer)
	}

	m.observers[name] = o
}

func (m *Message) RemoveObserver(name string) {
	delete(m.observers, name)
}

func (m *Message) NotifyObservers() {
	for _, v := range m.observers {
		v.Notify(m.Message)
	}
}
