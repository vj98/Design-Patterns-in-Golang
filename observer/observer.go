package observer

import "fmt"

type Observer interface {
	Update(message string)
}

type Subscriber struct {
	Name string
}

func (s *Subscriber) Update(message string) {
	fmt.Printf("%s received: %s\n", s.Name, message)
}

type Publisher struct {
	subscribers []Observer
}

func (p *Publisher) Subscribe(o Observer) {
	p.subscribers = append(p.subscribers, o)
}

func (p *Publisher) Notify(message string) {
	for _, sub := range p.subscribers {
		sub.Update(message)
	}
}
