package list

import "fmt"

// Bullet contains a rune
// is a representation
type Bullet struct {
	bullet rune
}

// NewBullet creates a new Bullet
func NewBullet(b rune) *Bullet {
	return &Bullet{
		bullet: b,
	}
}

// Print shows the content of the list
func (b *Bullet) Print(todos []string) {
	for _, v := range todos {
		fmt.Println("\t", string(b.bullet), v)
	}
}
