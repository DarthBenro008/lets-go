package book

import "newsfeeder/pkg/entities"

type Service interface {
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) InsertBook(book *entities.Book) (*entities.Book, error) {
	return s.InsertBook(book)
}
func (s *service) FetchBooks() (*[]entities.Book, error) {
	return s.FetchBooks()

}
func (s *service) UpdateBook(book *entities.Book) (*entities.Book, error) {
	return s.UpdateBook(book)
}
func (s *service) RemoveBook(ID string) error {
	return s.RemoveBook(ID)
}
