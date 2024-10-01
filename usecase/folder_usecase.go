package usecase

import "github.com/kurniawanxyz/crud-notes-go/domain"

type FolderUseCase struct{
	repo domain.FolderRepository
}

func NewFolderUseCase(repo domain.FolderRepository) *FolderUseCase{
	return &FolderUseCase{repo: repo}
}

func (uc *FolderUseCase) Index(UserID int) ([]domain.Folder, error){
	return uc.repo.Index(UserID)
}

func (uc *FolderUseCase) Show(UserID, ID int) (domain.Folder, error){
	return uc.repo.Show(UserID, ID)
}

func (uc *FolderUseCase) Store(folder domain.Folder) (*domain.Folder, error){
	return uc.repo.Store(&folder)
}

func (uc *FolderUseCase) Update(UserID, ID int, data *domain.UpdateFolderRequest) (*domain.Folder, error){
	return uc.repo.Update(UserID, ID, data)
}

func (uc *FolderUseCase) Delete(UserID, ID int) error{
	return uc.repo.Delete(UserID, ID)
}