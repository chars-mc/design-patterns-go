package main

import (
	"log"

	"github.com/chars-mc/design-patterns-go/structural/facade/facade"
)

func main() {
	f := facade.New("carlos@mail.com", "New comment", "valid-token", "user-blog")
	if err := f.Comment(); err != nil {
		log.Fatal(err)
	}
}
