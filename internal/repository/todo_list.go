package repository

import (
	"context"
	"github.com/diianpro/simpleApp/internal/models"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

type TodoList interface {
	Create(ctx context.Context, note *models.TodoList) error
	GetById(ctx context.Context, id uuid.UUID) (*models.TodoList, error)
	GetAll(ctx context.Context) ([]models.TodoList, error)
	Update(ctx context.Context, list *models.TodoList) error
	Delete(ctx context.Context, ID uuid.UUID) error
}

type todoList struct {
	db *pgxpool.Pool
}

func NewTodoList(db *pgxpool.Pool) *todoList {
	return &todoList{db: db}
}

func (t todoList) Create(ctx context.Context, note *models.TodoList) error {
	err := t.db.QueryRow(ctx, `INSERT into note(id, title, description) VALUES (gen_random_uuid(),$1,$2) RETURNING id`, note.Title, note.Description).Scan(&note.ID)
	return err
}

func (t todoList) GetById(ctx context.Context, ID uuid.UUID) (*models.TodoList, error) {
	row := t.db.QueryRow(ctx, `SELECT id,title, description from note where id = $1`, ID)
	return t.getById(row)
}

func (t todoList) getById(row pgx.Row) (*models.TodoList, error) {
	note := &models.TodoList{}
	err := row.Scan(&note.ID, &note.Title, &note.Description)
	if err != nil {
		log.Errorf("note db: getById error: %v", err)
		return nil, err
	}
	return note, err
}

func (t todoList) GetAll(ctx context.Context) ([]models.TodoList, error) {
	row, err := t.db.Query(ctx, `SELECT id, title, description FROM note`)
	if err != nil {
		log.Errorf("notes db: getAll error: %v", err)
		return nil, err
	}
	defer row.Close()

	var result []models.TodoList

	for row.Next() {

		newNote, err := t.getById(row)
		if err != nil {
			log.Errorf("notes db get data: getAll error: %v", err)
			return nil, err
		}

		result = append(result, *newNote)
	}

	return result, err
}

func (t todoList) Delete(ctx context.Context, ID uuid.UUID) error {
	_, err := t.db.Exec(ctx, `DELETE from note where id = $1`, ID)
	return err
}

func (t todoList) Update(ctx context.Context, list *models.TodoList) error {
	_, err := t.db.Query(ctx, `UPDATE note SET  title = $1,  description = $2  WHERE id = $3`, list.Title, list.Description, list.ID)
	return err
}
