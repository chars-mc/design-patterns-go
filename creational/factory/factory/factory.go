package factory

import "github.com/chars-mc/design-patterns-go/creational/factory/connection"

// Factory returns a DBConnection interface
func Factory(t int) connection.DBConnection {
	switch t {
	case 1:
		return &connection.MySQL{}
	case 2:
		return &connection.PostgreSQL{}
	default:
		return nil
	}
}
