package database

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	for tries := 0; tries < 3; tries++ {
		conn, err := sql.Open("mysql", os.Getenv("DATABASE_URI"))
		if err != nil {
			if tries == 2 {
				log.Fatalf("Error connecting to database: %v", err)
			}
			continue
		}

		conn.SetConnMaxIdleTime(time.Minute * 3)
		conn.SetMaxOpenConns(10)
		conn.SetMaxIdleConns(10)

		DB = conn
		break
	}
}
