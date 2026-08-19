package repo
import "myapp/models"

type ArticleRepository struct {
	articles []models.Article
}

func NewArticleRepository() *ArticleRepository {
	return &ArticleRepository{
		articles: []models.Article{
			{
				ID: 1,
				Title: "Hello",
				Desc: "Article Description",
				Content: "Article Content",
			},
			{
				ID: 2,
				Title: "Hello 1",
				Desc: "Article Description 1",
				Content: "Article Content 1",
			},
		},
	}
}

func (r *ArticleRepository) GetAll() []models.Article {
	return r.articles
}

func (r *ArticleRepository) GetByID(id int) *models.Article {
	for _, article := range r.articles {
		if article.ID == id {
			return &article
		}
	}

	return nil
}
