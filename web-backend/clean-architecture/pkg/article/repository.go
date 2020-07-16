package article


type ArticleRepository interface {
}

type articleRepo struct {
	Articles []Articles
}

func NewArticleRepo() *articleRepo {
	return &articleRepo{
		Articles: []Articles{},
	}
}

func (r *articleRepo) checkArticleByArticle(article Article) bool {
	for _, presentArticle := range r.Articles {
		if presentArticle.Title == article.Title {
			return true
		}
	}
	return false
}


func (r *articleRepo) checkArticleById(Id int) bool {
	for _, presentArticle := range r.Articles {
		if presentArticle.Id == article.Id {
			return true
		}
	}
	return false
}

func (r *articleRepo) upsertArticle(article Article) bool{
	r.Articles = append(r.Articles, article)
	return true
}

func (r *articleRepo) updateArticle(article Article) bool {
	for index, presentArticle := range r.Articles{
		if presentArticle.Title == article.Title{
			r.Articles[index] = article
		}
	}
	return false
}

func (r *articleRepo) deleteArticle(id int) bool{
	for index,presentArticle := range r.Items{
		if presentArticle.Id == article.Id{
			frontList := r.Articles[:index]
			backList := r.Articles[index+1:]
			r.Articles = append(frontList,backList)
			return true
		}
	}
	return false
}