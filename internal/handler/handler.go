package handler

import (
	_ "github.com/diianpro/simpleApp/docs"
	"github.com/diianpro/simpleApp/internal/models"
	"github.com/diianpro/simpleApp/internal/service"
	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

type NoteHandler struct {
	noteHandler *service.NoteService
}

//CreateNote godoc
//@Summary Create note based on given ID
//@Tags note
//@Produce json
//@Param title path string true "Title"
//@Param description path string true "Description"
//@Success 200  {object} models.TodoList
//@Router /note [post]
func (h *NoteHandler) CreateNote(c echo.Context) error {
	var request models.TodoList
	req := request
	if err := c.Bind(&req); err != nil {
		log.Errorf("Create: Bind error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.noteHandler.Create(c.Request().Context(), &req)
	if err != nil {
		log.Errorf("Create: Note error: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, req.ID)
}

//DeleteNote godoc
//@Summary Delete note based on given ID
//@Tags note
//@Produce json
//@Param id path int true "Note ID"
//@Success 200 {object} models.TodoList
//@Router /note/:id [delete]
func (h *NoteHandler) DeleteNote(c echo.Context) error {
	type request struct {
		ID uuid.UUID `param:"id"`
	}
	req := request{}
	if err := c.Bind(&req); err != nil {
		log.Errorf("Delete: Bind error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err := h.noteHandler.Delete(c.Request().Context(), req.ID)
	if err != nil {
		log.Errorf("Delete: Note error: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

//GetByID godoc
//@Summary Retrieves note based on given ID
//@Tags note
//@Produce json
//@Param id path int true "Note ID"
//@Success 200 {object} models.TodoList
//@Router /note/:id [get]
func (h *NoteHandler) GetByID(c echo.Context) error {
	type request struct {
		ID uuid.UUID `param:"id"`
	}
	req := request{}
	if err := c.Bind(&req); err != nil {
		log.Errorf("GetByID: Bind error: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	notes, err := h.noteHandler.GetById(c.Request().Context(), req.ID)
	if err != nil {
		log.Errorf("GetByID: Note error: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, notes)
}

//GetAll godoc
//@Summary GetAll note
//@Tags note
//@Produce json
//@Success 200 {array} models.TodoList
//@Router /notes [get]
func (h *NoteHandler) GetAll(c echo.Context) error {
	notes, err := h.noteHandler.GetAll(c.Request().Context())
	if err != nil {
		log.Errorf("GetAll: Notes error: %v", err)
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, notes)
}

//UpdateNote godoc
//@Summary UpdateNote note based on given ID
//@Tags note
//@Produce json
//@Param title path string true "Title"
//@Param description path string true "Description"
//@Success 200 {object} models.TodoList
//@Router /note/:id [put]
func (h *NoteHandler) UpdateNote(c echo.Context) error {
	var request *models.TodoList
	req := request
	if err := c.Bind(&req); err != nil {
		log.Errorf("Update: Note error: %v", err)
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	err := h.noteHandler.Update(c.Request().Context(), req)
	if err != nil {
		log.Errorf("Update: Note error: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, req)
}

func NewNoteHandler(noteHandler *service.NoteService) *NoteHandler {
	return &NoteHandler{noteHandler: noteHandler}
}
