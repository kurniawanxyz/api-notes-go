package usecase

import "github.com/kurniawanxyz/crud-notes-go/domain"

type NoteUseCase struct {
	repo domain.NoteRepository
}

func NewNoteUseCase(repo domain.NoteRepository) *NoteUseCase {
	return &NoteUseCase{repo}
}

func (u *NoteUseCase) Index(userID, folderID int) ([]domain.Note, error) {
	return u.repo.Index(userID, folderID)
}

func (u *NoteUseCase) Show(userID, folderID, id int) (domain.Note, error) {
	return u.repo.Show(userID, folderID, id)
}

func (u *NoteUseCase) Store(note *domain.Note) (*domain.Note, error) {
	return u.repo.Store(note)
}

func (u *NoteUseCase) Update(userID, folderID, id int, data *domain.Note) (*domain.Note, error) {
	return u.repo.Update(userID, folderID, id, data)
}

func (u *NoteUseCase) Delete(userID, folderID, id int) error {
	return u.repo.Delete(userID, folderID, id)
}


