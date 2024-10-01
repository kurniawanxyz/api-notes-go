package repository

import (
	"github.com/kurniawanxyz/crud-notes-go/domain"
	"gorm.io/gorm"
)

type NoteRepository struct{
	DB *gorm.DB
}

func NewNoteRepository (db *gorm.DB) *NoteRepository {
	return &NoteRepository{DB: db}
}

func (r *NoteRepository) Index(UserID, FolderID int) (data []domain.Note,err error) {
	tx := r.DB.Begin()
	if err := r.DB.Find(&data, "user_id = ? AND folder_id = ?", UserID, FolderID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return data, nil
}

func (r *NoteRepository) Show(UserID, FolderID, ID int) (data domain.Note, err error){
	tx := r.DB.Begin()

	if err := r.DB.First(&data, "user_id = ? AND folder_id = ? AND id = ?", UserID, FolderID, ID).Error; err != nil {
		tx.Rollback()
		return data, err
	}

	tx.Commit()
	return data, nil
}

func (r *NoteRepository) Store(data *domain.Note) (result *domain.Note, err error){
	tx := r.DB.Begin()

	if err := r.DB.Create(&data).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return data, nil
}

func (r *NoteRepository) Update(UserID, FolderID, ID int, data *domain.Note) (result *domain.Note, err error){
	tx := r.DB.Begin()
	if err := r.DB.First(&data, "user_id = ? AND folder_id = ? AND id = ? ", UserID,FolderID,ID).Save(data).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return data, nil
}

func (r *NoteRepository) Delete(UserID, FolderID, ID int) (err error){
	tx := r.DB.Begin()

	if err := r.DB.Where("user_id = ? AND folder_id = ? AND id = ?", UserID, FolderID, ID).Delete(&domain.Note{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}