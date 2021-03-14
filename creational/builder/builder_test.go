package builder_test

import (
	"testing"

	"github.com/chars-mc/design-patterns-go/creational/builder"
)

func TestSender_BuildMessage(t *testing.T) {
	json := &builder.JSONMessageBuilder{}
	sender := &builder.Sender{}

	sender.SetBuilder(json)
	msg, err := sender.BuildMessage("Gopher", "hello world")
	if err != nil {
		t.Fatalf("BuildMessage(): unespected error, got %v", err)
	}
	t.Log(string(msg.Body))
}
