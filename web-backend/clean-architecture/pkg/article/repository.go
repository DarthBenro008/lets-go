package article

import "articlefeeder/pkg/entities"

//Repository holds all the key functions related to the entity
type Repository interface {
	CheckArticleByArticle(article entities.Article) bool
	CheckArticleByID(ID int) bool
	UpsertArticle(article entities.Article) bool
	UpdateArticle(article entities.Article) bool
	DeleteArticle(id int) bool
	FetchAllArticles() []entities.Article
}

type articleRepo struct {
	Articles []entities.Article
}

//NewArticleRepo creates a new instance of this repo
func NewArticleRepo() *articleRepo {
	return &articleRepo{
		Articles: []entities.Article{},
	}
}

func (r *articleRepo) CheckArticleByArticle(article entities.Article) bool {
	for _, presentArticle := range r.Articles {
		if presentArticle.Title == article.Title {
			return true
		}
	}
	return false
}

func (r *articleRepo) CheckArticleByID(ID int) bool {
	for _, presentArticle := range r.Articles {
		if presentArticle.ID == ID {
			return true
		}
	}
	return false
}

func (r *articleRepo) UpsertArticle(article entities.Article) bool {
	r.Articles = append(r.Articles, article)
	return true
}

func (r *articleRepo) UpdateArticle(article entities.Article) bool {
	for index, presentArticle := range r.Articles {
		if presentArticle.Title == article.Title {
			r.Articles[index] = article
		}
	}
	return false
}

func (r *articleRepo) DeleteArticle(id int) bool {
	for index, presentArticle := range r.Articles {
		if presentArticle.ID == id {
			frontList := r.Articles[:index]
			backList := r.Articles[index+1:]
			r.Articles = append(frontList, backList...)
			return true
		}
	}
	return false
}

func (r *articleRepo) FetchAllArticles() []entities.Article {
	return r.Articles
}
