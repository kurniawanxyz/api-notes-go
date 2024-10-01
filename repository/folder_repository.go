package repository

import (
	"github.com/kurniawanxyz/crud-notes-go/domain"
	"gorm.io/gorm"
)

type FolderRepository struct{
	DB *gorm.DB
}

func NewFolderRepository(db *gorm.DB) *FolderRepository {
	return &FolderRepository{DB: db}
}

func (r *FolderRepository) Index(UserId int) (folders []domain.Folder,err error){
	if err = r.DB.Find(&folders, "user_id = ?", UserId).Error; err != nil {
		return nil, err
	}
	return folders, nil
}

func (r *FolderRepository) Show(UserId, ID int) (folder domain.Folder, err error){
	if err = r.DB.First(&folder, "id = ? AND user_id = ?",ID, UserId).Error; err != nil {
		return domain.Folder{}, err
	}
	return folder, nil
}

func (r *FolderRepository) Store(folder *domain.Folder) (*domain.Folder, error){
	if err := r.DB.Create(folder).Error; err != nil {
		return nil, err
	}
	return folder, nil
}

func (r *FolderRepository) Update(UserId, ID int, data *domain.UpdateFolderRequest) (*domain.Folder, error){
	var folder domain.Folder
	if err := r.DB.First(&folder,"id = ? AND user_id = ?",ID,UserId).Save(data).Error; err != nil {
		return nil, err
	}
	return &folder, nil
}

func (r *FolderRepository) Delete(UserId, ID int) error{
	if err := r.DB.Where("id = ? AND user_id = ?", ID, UserId).Delete(&domain.Folder{}).Error; err != nil {
		return err
	}
	return nil
}