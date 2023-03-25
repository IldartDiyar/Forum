package app

import (
	"log"
	"net/http"

	"forum_backend/internal/Log"
	"forum_backend/internal/database"
	"forum_backend/internal/handler"
	"forum_backend/internal/repository"
	"forum_backend/internal/service"
)

func Run() error {
	logger, err := Log.NewFileLogger()
	if err != nil {
		log.Fatal(err)
	}
	logger.Debug("Successfully Initiated the Loger")

	configDb := database.NewConfDb(logger)
	db := configDb.InitDB(logger)
	defer db.Close()
	logger.Debug("Successfully Initiated the Data Base")

	configDb.CreateTables(db, logger)
	logger.Debug("Successfully created the tables")

	repo := repository.NewRepository(db)
	logger.Debug("Successfully Initiated the Repository")

	service := service.NewService(repo)
	logger.Debug("Successfully Initiated the Service")

	handler := handler.NewHandler(service, logger)
	return http.ListenAndServe(":8080", handler.Start())
}
