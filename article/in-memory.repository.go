package article

import (
	"errors"
)

type inMemoryArticleRepository struct {
	articles []Article
}

var inMemoArticleRepo *inMemoryArticleRepository = nil


func InMemoryArticleRepository() ArticleRepository {
	if inMemoArticleRepo == nil {
		inMemoArticleRepo = new(inMemoryArticleRepository)
	}
	return inMemoArticleRepo
}

func (r *inMemoryArticleRepository) Create(article *Article) error {
	r.articles = append(r.articles, *article)
	return nil
}

func (r *inMemoryArticleRepository) Update(input *Article) error {
	for i:=0; i<len(r.articles); i++ {
		if r.articles[i].Id == input.Id {
			r.articles[i].Body = input.Body
			r.articles[i].Description = input.Description
			r.articles[i].Title = input.Title
			r.articles[i].Slug = input.Slug
			return nil
		}
	}
	return errors.New("Article does not exist")
}

func (r *inMemoryArticleRepository) FindBySlug(slug string) *Article {
	article := new(Article)
	for _, v := range r.articles{
		if v.Slug == slug {
			*article = v
			return article
		}
	}
	return nil
}

func (r *inMemoryArticleRepository) FindAll() []Article {
	retval := make([]Article, len(r.articles))
	copy(retval, r.articles)

	return retval
}