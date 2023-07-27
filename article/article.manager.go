package article

import (
	"realworld-backend/shared"

	"github.com/gosimple/slug"
)

func CreateArticle(repo ArticleRepository, idGenerator shared.IDGenerator) func (*ArticleInput) (*Article, error) {
	return func (input *ArticleInput) (*Article, error) {
		id := idGenerator()
		article := new(Article)

		article.Body = input.Body
		article.Title = input.Title
		article.Description = input.Description
		article.Id = id
		article.Slug = slug.Make(article.Title)

		err := repo.Create(article)
		if (err != nil) {
			return nil, err
		} else {
			return article, nil
		}
	}
}