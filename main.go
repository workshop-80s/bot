package main

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestMySQLConnection(t *testing.T) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		t.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		t.Fatalf("Error pinging database: %v", err)
	}

	// Test table creation and query
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS test (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255))")
	if err != nil {
		t.Fatalf("Error creating table: %v", err)
	}

	_, err = db.Exec("INSERT INTO test (name) VALUES (?)", "circleci-test")
	if err != nil {
		t.Fatalf("Error inserting data: %v", err)
	}

	var name string
	err = db.QueryRow("SELECT name FROM test WHERE name = ?", "circleci-test").Scan(&name)
	if err != nil {
		t.Fatalf("Error querying data: %v", err)
	}

	if name != "circleci-test" {
		t.Errorf("Expected name to be 'circleci-test', got '%s'", name)
	}
}
