package domain

import (
	"time"
)

type Note struct {
	ID        int       `json:"id" gorm:"primary_key"`
	UserID    int       `json:"user_id" gorm:"index" validate:"required"`
	FolderID  int       `json:"folder_id" gorm:"index" validate:"required"`
	Title     string    `json:"title" validate:"required"`
	Content   string    `json:"content" validate:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}

type NoteRepository interface {
	Index(userID, folderID int) ([]Note, error)
	Show(userID, folderID, id int) (Note, error)
	Store(note *Note) (*Note, error)
	Update(userID, folderID, id int, data *Note) (*Note, error)
	Delete(userID, folderID, id int) error
}
