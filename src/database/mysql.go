package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

func createMysqlConnectionInfo(host string, port int, user string, pass string, dbName string) string {
	info := fmt.Sprintf("%s:%s@(%s:%d)/%s", user, pass, host, port, dbName)
	info += "?charset=utf8&parseTime=True&loc=UTC"
	return info
}

func OpenMySql() (*sql.DB, error) {
	connection := createMysqlConnectionInfo("127.0.0.1", 3306, "user", "password", "test_docker_db")
	return sql.Open("mysql", connection)
}

func Ping(db *sql.DB, ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
}
