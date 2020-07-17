package book

import "newsfeeder/pkg/entities"

type Service interface {
	InsertBook(book *entities.Book) (*entities.Book, error)
	FetchBooks() (*[]entities.Book, error)
	UpdateBook(book *entities.Book) (*entities.Book, error)
	RemoveBook(ID string) error
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
