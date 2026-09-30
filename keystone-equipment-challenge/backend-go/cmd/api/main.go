package main

import (
	"log"
	"net/http"

	"keystone/equipment-api/internal/config"
	"keystone/equipment-api/internal/db"
	"keystone/equipment-api/internal/server"
)

func main() {
	config.LoadEnv()
	db.Open()

	port := config.Get("PORT", "8080")
	log.Printf("Keystone equipment API (Go) listening on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, server.New()))
}
