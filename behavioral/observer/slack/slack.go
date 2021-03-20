package slack

import "fmt"

type Slack struct{}

func (s *Slack) Notify(data string) {
	sendMessage(data)
}

func sendMessage(data string) {
	fmt.Printf("message sent to slack: %q", data)
}
