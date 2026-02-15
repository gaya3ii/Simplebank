package db

import (
	"bufio"
	"database/sql"
	"os"
	"path/filepath"
	"strings"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var dbURL = os.Getenv("DATABASE_URL")

var testStore *Queries

// loadEnv reads .env from the project root and sets env vars (e.g. DATABASE_URL).
func loadEnv() {
	dir, _ := os.Getwd()
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			break
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return
		}
		dir = parent
	}
	f, err := os.Open(filepath.Join(dir, ".env"))
	if err != nil {
		return
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		line = strings.TrimPrefix(line, "export ")
		if i := strings.Index(line, "="); i > 0 {
			key := strings.TrimSpace(line[:i])
			val := strings.TrimSpace(line[i+1:])
			val = strings.Trim(val, `"`)
			os.Setenv(key, val)
		}
	}
}

func skipIfNoDB(t *testing.T) {
	if testStore == nil {
		t.Skip("no database: set DATABASE_URL and ensure DB exists to run integration tests")
	}
}

func TestMain(m *testing.M) {
	loadEnv()
	dbURL = os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/simplebank?sslmode=disable"
	}
	conn, err := sql.Open("pgx", dbURL)
	if err != nil {
		os.Exit(m.Run())
		return
	}
	if err := conn.Ping(); err != nil {
		conn.Close()
		os.Exit(m.Run())
		return
	}
	testStore = New(conn)
	os.Exit(m.Run())
}
