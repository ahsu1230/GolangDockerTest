package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func OpenMySql(host string, port int, user string, pass string, dbName string) (*sql.DB, error) {
	connection := fmt.Sprintf("%s:%s@(%s:%d)/%s", user, pass, host, port, dbName)
	fmt.Println("Opening connection to", connection)
	return sql.Open("mysql", connection)
}

func Ping(db *sql.DB, ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return err
	}
	return nil
}
