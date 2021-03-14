package main

import (
	"fmt"
	"log"

	"github.com/chars-mc/design-patterns-go/creational/factory/factory"
)

func main() {
	var t int
	fmt.Println("Database type: ")
	_, err := fmt.Scan(&t)
	if err != nil {
		log.Fatal(err)
	}

	conn := factory.Factory(1)
	if conn == nil {
		log.Fatal("engine not valid")
	}

	err = conn.Connect()
	if err != nil {
		log.Fatal("cannot connect to db server", err)
	}

	now, err := conn.GetNow()
	if err != nil {
		log.Fatal("cannot get system time")
	}

	fmt.Println(now)

	err = conn.Close()
	if err != nil {
		log.Fatal(err)
	}
}
