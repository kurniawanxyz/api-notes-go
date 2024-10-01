package domain

import "time"

type Folder struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	UserID      int       `json:"user_id" gorm:"foreignKey:UserID"`
	Name        string    `json:"name" gorm:"required" validate:"required,min=3,max=255"`
	Description string    `json:"description" gorm:"required" validate:"required,min=3,max=255"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}

type UpdateFolderRequest struct {
	Name        string `json:"name" validate:"min=3,max=255"`
	Description string `json:"description" validate:"min=3,max=255"`
}

type FolderRepository interface {
	Index(userID int) ([]Folder, error)
	Show(userID, id int) (Folder, error)
	Store(folder *Folder) (*Folder, error)
	Update(userID, id int, data *UpdateFolderRequest) (*Folder, error)
	Delete(userID, id int) error
}
