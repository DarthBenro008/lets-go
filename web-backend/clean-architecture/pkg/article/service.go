package article

import (
	"articlefeeder/pkg/entities"
)

type ArticleService interface {
	InsertArticle(article entities.Article) bool
	UpdateArticle(article entities.Article) bool
	DeleteArticle(id int) bool
	FetchAll() []entities.Article
}

type articleService struct {
	serviceRepo Repository
}

func NewService(r Repository) ArticleService {
	return &articleService{
		serviceRepo: r,
	}
}

func (s *articleService) InsertArticle(article entities.Article) bool {
	if g := s.serviceRepo.CheckArticleByArticle(article); g {
		return s.serviceRepo.UpsertArticle(article)
	}
	return false
}

func (s *articleService) UpdateArticle(article entities.Article) bool {
	return s.serviceRepo.UpdateArticle(article)
}
func (s *articleService) DeleteArticle(id int) bool {
	if g := s.serviceRepo.CheckArticleById(id); g {
		return s.serviceRepo.DeleteArticle(id)
	}
	return false
}
func (s *articleService) FetchAll() []entities.Article {
	return s.serviceRepo.FetchAllArticles()
}
