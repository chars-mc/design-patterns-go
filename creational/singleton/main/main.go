package main

import (
	"fmt"
	"sync"

	"github.com/chars-mc/design-patterns-go/creational/singleton"
	client_one "github.com/chars-mc/design-patterns-go/creational/singleton/client-one"
	client_two "github.com/chars-mc/design-patterns-go/creational/singleton/client-two"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(200)

	client_one.SetAge(3)
	client_two.SetAge(6)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			go client_one.IncrementAge()
		}()
		go func() {
			defer wg.Done()
			go client_two.IncrementAge()
		}()
	}

	wg.Wait()
	p := singleton.GetInstance()
	fmt.Println(p.GetAge())
}
