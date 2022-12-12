package main

import (
	"context"
	_ "github.com/diianpro/simpleApp/docs"
	"github.com/diianpro/simpleApp/internal/config"
	"github.com/diianpro/simpleApp/internal/handler"
	"github.com/diianpro/simpleApp/internal/repository"
	"github.com/diianpro/simpleApp/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
)

// @title Swagger RestAPI
// @version 2.0
// @description This is a sample server RestAPI server.

// @host localhost:1323
// @BasePath /

func main() {
	ctx := context.Background()
	cfg, err := config.NewConfig(ctx)
	if err != nil {
	}
	noteRps := repository.NewTodoList(cfg.PgConnection)
	noteService := service.NewNoteService(noteRps)
	noteHandler := handler.NewNoteHandler(noteService)
	e := echo.New()

	e.POST("/note", noteHandler.CreateNote)
	e.GET("/note/:id", noteHandler.GetByID)
	e.GET("/notes", noteHandler.GetAll)
	e.DELETE("/note/:id", noteHandler.DeleteNote)
	e.PUT("/note/:id", noteHandler.UpdateNote)

	e.GET("/docs/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
