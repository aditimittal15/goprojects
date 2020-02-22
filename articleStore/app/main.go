package main

import (
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

const (
	defaultServerPort = "8080"
)

type ServiceCgf struct {
	Port string
}

var cfg ServiceCgf

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func readEnvConfiguration() {
	cfg.Port = getenv("PORT", defaultServerPort)
}
func main() {
	readEnvConfiguration()
	router := NewRouter()
	server := &http.Server{
		Handler:      router,
		Addr:         ":" + cfg.Port,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
