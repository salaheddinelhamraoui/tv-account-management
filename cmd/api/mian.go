// Package main starts the API service.
package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/salaheddinelhamraoui/tv-account-management/internal/config"
	"github.com/salaheddinelhamraoui/tv-account-management/internal/database"
	"github.com/salaheddinelhamraoui/tv-account-management/internal/logger"
)

func main() {
	log := logger.New()
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}

	db, err := database.New(cfg.Database)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
	}

	mainDB, err := db.DB()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to main database")
	}

	defer func(mainDB *sql.DB) {
		err := mainDB.Close()
		if err != nil {
			fmt.Printf("failed to close main database: %v\n", err)
		}
	}(mainDB)
	gin.SetMode(cfg.Server.GinMode)

	log.Info().Msg("starting server")
}
