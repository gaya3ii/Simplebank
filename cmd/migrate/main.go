// One-off: run db/migration/000001_init_schema.up.sql.
// Usage: source .env && go run ./cmd/migrate
package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	connURL := os.Getenv("DATABASE_URL")
	if connURL == "" {
		fmt.Fprintln(os.Stderr, "DATABASE_URL is not set (e.g. source .env)")
		os.Exit(1)
	}
	conn, err := sql.Open("pgx", connURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "connect: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()
	if err := conn.Ping(); err != nil {
		fmt.Fprintf(os.Stderr, "ping: %v\n", err)
		os.Exit(1)
	}
	sqlBytes, err := os.ReadFile("db/migration/000001_init_schema.up.sql")
	if err != nil {
		fmt.Fprintf(os.Stderr, "read migration: %v\n", err)
		os.Exit(1)
	}
	for i, stmt := range strings.Split(string(sqlBytes), ";") {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}
		if _, err := conn.Exec(stmt + ";"); err != nil {
			fmt.Fprintf(os.Stderr, "statement %d: %v\n", i+1, err)
			os.Exit(1)
		}
	}
	fmt.Println("Migration applied.")
}
