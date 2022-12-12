package repository

import (
	"context"
	"github.com/diianpro/simpleApp/internal/config"
	"github.com/diianpro/simpleApp/internal/handler"
	"github.com/diianpro/simpleApp/internal/service"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

var db *pgxpool.Pool

func TestMain(m *testing.M) {
	ctx := context.Background()
	cfg, err := config.NewConfig(ctx)
	if err != nil {
		log.Errorf("error in config: initialize store config: %v", err)
		return
	}
	noteRps := NewTodoList(cfg.PgConnection)
	noteService := service.NewNoteService(noteRps)
	noteHandler := handler.NewNoteHandler(noteService)

	e := echo.New()
	defer func() {
		err := e.Shutdown(ctx)
		if err != nil {
			log.Errorf("Error: %v", err)
		}
	}()

	e.POST("/note", noteHandler.CreateNote)
	e.GET("/note/:id", noteHandler.GetByID)
	e.GET("/notes", noteHandler.GetAll)
	e.DELETE("/note/:id", noteHandler.DeleteNote)
	e.PUT("/note/:id", noteHandler.UpdateNote)

	go e.Logger.Fatal(e.Start(":1323"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	m.Run()

}
