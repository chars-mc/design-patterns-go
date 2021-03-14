package connection

import "time"

// DBConnection contains the methods to interact with any database engine
type DBConnection interface {
	Connect() error
	GetNow() (*time.Time, error)
	Close() error
}
