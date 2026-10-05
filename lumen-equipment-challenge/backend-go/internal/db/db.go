package db

import (
	"database/sql"
	"log"
	"path/filepath"

	"keystone/equipment-api/internal/config"
)

var DB *sql.DB

func Open() {
	path := filepath.Join(config.RepoRoot(), config.Get("DB_PATH", "db/keystone.db"))
	var err error
	DB, err = sql.Open(driverName, path)
	if err != nil {
		log.Fatal(err)
	}
}
