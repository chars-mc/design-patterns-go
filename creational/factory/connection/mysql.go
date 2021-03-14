package connection

import (
	"database/sql"
	"fmt"
	"time"

	// driver to connect to MySQL
	_ "github.com/go-sql-driver/mysql"
)

// MySQL contains the db connection
type MySQL struct {
	db *sql.DB
}

// Connect connects to the db server
func (m *MySQL) Connect() error {
	// user:password@protocol(server:port)/database?params...
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?tls=false&autocommit=true&allowNativePasswords=true&parseTime=true",
		"root",
		"userRootPass",
		"127.0.0.1",
		"3306",
		"store",
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}

	m.db = db
	return nil
}

// GetNow returns the current date of the db
func (m *MySQL) GetNow() (*time.Time, error) {
	t := &time.Time{}
	err := m.db.QueryRow("SELECT CURDATE()").Scan(t)
	if err != nil {
		fmt.Printf("error getting the current date of the server %v", err)
		return nil, err
	}
	return t, nil
}

// Close closes the db connection
func (m *MySQL) Close() error {
	return m.db.Close()
}
