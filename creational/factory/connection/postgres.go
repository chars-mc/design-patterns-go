package connection

import (
	"database/sql"
	"fmt"
	"time"
)

// PostgreSQL contains the db connection
type PostgreSQL struct {
	db *sql.DB
}

// Connect connects to the db server
func (p *PostgreSQL) Connect() error {
	// postgres://username:password@server:port/database?params...
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		"postgres",
		"postgres",
		"127.0.0.1",
		"5432",
		"postgres",
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	p.db = db
	return nil
}

// GetNow returns the current date of the db
func (p *PostgreSQL) GetNow() (*time.Time, error) {
	t := &time.Time{}
	err := p.db.QueryRow("SELECT CURRENT_DATE").Scan(t)
	if err != nil {
		fmt.Printf("error getting the current date of the server %v", err)
		return nil, err
	}
	return t, nil
}

// Close closes the db connection
func (p *PostgreSQL) Close() error {
	return p.db.Close()
}
