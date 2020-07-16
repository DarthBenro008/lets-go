package article

import (
	"articlefeeder/pkg/entities"
)

type ArticleService interface {
	InsertArticle(article entities.Article) bool
	UpdateArticle(article entities.Article) bool
	DeleteArticle(id int) bool
	fetchAll() []entities.Article
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
	if g := s.serviceRepo.checkArticleByArticle(article); g {
		return s.serviceRepo.upsertArticle(article)
	}
	return false
}

func (s *articleService) UpdateArticle(article entities.Article) bool {
	return s.serviceRepo.updateArticle(article)
}
func (s *articleService) DeleteArticle(id int) bool {
	if g := s.serviceRepo.checkArticleById(id); g {
		return s.serviceRepo.deleteArticle(id)
	}
	return false
}
func (s *articleService) fetchAll() []entities.Article {
	return s.serviceRepo.fetchAllArticles()
}
