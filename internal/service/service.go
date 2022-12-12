package service

import (
	"context"
	"github.com/diianpro/simpleApp/internal/models"
	"github.com/diianpro/simpleApp/internal/repository"
	"github.com/gofrs/uuid"
)

type NoteService struct {
	noteRps repository.TodoList
}

func (n *NoteService) Create(ctx context.Context, note *models.TodoList) error {
	return n.noteRps.Create(ctx, note)
}

func (n *NoteService) Delete(ctx context.Context, ID uuid.UUID) error {
	return n.noteRps.Delete(ctx, ID)
}

func (n *NoteService) GetById(ctx context.Context, id uuid.UUID) (*models.TodoList, error) {
	return n.noteRps.GetById(ctx, id)
}

func (n *NoteService) GetAll(ctx context.Context) ([]models.TodoList, error) {
	return n.noteRps.GetAll(ctx)
}

func (n *NoteService) Update(ctx context.Context, list *models.TodoList) error {
	return n.noteRps.Update(ctx, list)
}

func NewNoteService(noteRps repository.TodoList) *NoteService {
	return &NoteService{
		noteRps: noteRps,
	}
}
