package article

import "articlefeeder/pkg/entities"

//Repository holds all the key functions related to the entity
type Repository interface {
	checkArticleByArticle(article entities.Article) bool
	checkArticleById(ID int) bool
	upsertArticle(article entities.Article) bool
	updateArticle(article entities.Article) bool
	deleteArticle(id int) bool
	fetchAllArticles() []entities.Article
}

type articleRepo struct {
	Articles []entities.Article
}

func NewArticleRepo() *articleRepo {
	return &articleRepo{
		Articles: []entities.Article{},
	}
}

func (r *articleRepo) checkArticleByArticle(article entities.Article) bool {
	for _, presentArticle := range r.Articles {
		if presentArticle.Title == article.Title {
			return true
		}
	}
	return false
}

func (r *articleRepo) checkArticleById(Id int) bool {
	for _, presentArticle := range r.Articles {
		if presentArticle.ID == Id {
			return true
		}
	}
	return false
}

func (r *articleRepo) upsertArticle(article entities.Article) bool {
	r.Articles = append(r.Articles, article)
	return true
}

func (r *articleRepo) updateArticle(article entities.Article) bool {
	for index, presentArticle := range r.Articles {
		if presentArticle.Title == article.Title {
			r.Articles[index] = article
		}
	}
	return false
}

func (r *articleRepo) deleteArticle(id int) bool {
	for index, presentArticle := range r.Articles {
		if presentArticle.ID == id {
			frontList := r.Articles[:index]
			backList := r.Articles[index+1:]
			r.Articles = append(frontList, backList)
			return true
		}
	}
	return false
}

func (r *articleRepo) fetchAllArticles() []entities.Article {
	return r.Articles
}
