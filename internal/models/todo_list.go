package models

import "github.com/gofrs/uuid"

type TodoList struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}
