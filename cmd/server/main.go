package main

import (
	"Datearn/internal/config"
	"Datearn/internal/server/http"
	"Datearn/pkg/database"
	"flag"
	"fmt"
	"log"
)

// migration defines migration flag
var migration *bool = flag.Bool("migrate", false, "Run migrations")

// main Where every shit starts
func main() {
	flag.Parse()
	err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	if *migration {
		migrate()
		return
	}

	http.StartHttpServer()
}

// migrate runs and executed migrations for sql structure
func migrate() {
	err := database.Migrate()
	if err != nil {
		fmt.Printf("%s: %s", "Running migrations faild with error", err)
	}
}
