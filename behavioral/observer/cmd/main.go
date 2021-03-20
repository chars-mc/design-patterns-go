package main

import (
	"github.com/chars-mc/design-patterns-go/behavioral/observer/email"
	"github.com/chars-mc/design-patterns-go/behavioral/observer/interfaces"
	message "github.com/chars-mc/design-patterns-go/behavioral/observer/message"
	"github.com/chars-mc/design-patterns-go/behavioral/observer/slack"
)

func main() {
	m := message.Message{}
	m.AddObserver("slack", observerFactory("slack"))
	m.Message = "This is a new message"
	m.NotifyObservers()
}

func observerFactory(name string) interfaces.Observer {
	switch name {
	case "slack":
		return &slack.Slack{}
	case "email":
		return &email.Email{}
	}
	return nil
}
